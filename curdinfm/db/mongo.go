
type MongoQueryMultiArgs struct {
	TableName  string
	Conds      bson.M
	Cols       bson.M
	SortCol    bson.M
	DecodeType interface{} // the point of Decode Type
	Limit      int
}

// MongoQueryMulti 泛型 - mongo 查询集合
func MongoQueryMulti(tx *mgo.Database, args MongoQueryMultiArgs) (rst interface{}, err error) {
	defer CheckPanic(&err)

	outType := reflect.TypeOf(args.DecodeType)
	outs := reflect.MakeSlice(reflect.SliceOf(outType), 0, 0)

	limitTimeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), limitTimeout)
	defer cancel()

	opts := []*options.FindOptions{}
	opts = append(opts, options.Find().SetProjection(args.Cols))
	opts = append(opts, options.Find().SetSort(args.SortCol))
	if args.Limit != 0 {
		opts = append(opts, options.Find().SetLimit(int64(args.Limit)))
	}
	cursor, err := tx.Collection(args.TableName).Find(ctx, args.Conds, opts...)
	if err != nil && err.Error() != enum.MgoNotFoundErr {
		return nil, fmt.Errorf("mongo-driver collection.find failed %s", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		vals := reflect.New(outType.Elem()).Interface()

		if err = cursor.Decode(vals); err != nil {
			return nil, fmt.Errorf("mongo-driver Deserialize failed %s", err)
		}
		if err = cursor.Err(); err != nil {
			return nil, fmt.Errorf("mongo-driver cursor failed %s", err)
		}
		outs = reflect.Append(outs, reflect.ValueOf(vals))
	}
	rst = outs.Interface()
	return rst, nil
}

// CheckPanic  Panic --> Error
func CheckPanic(err *error) {
	p := recover()
	if p != nil {
		*err = fmt.Errorf("Catch Panic [%v]", p)
	}
}

