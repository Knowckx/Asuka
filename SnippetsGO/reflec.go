package SnippetsGO

import (
	"fmt"
	"reflect"
)

type MongoQueryMultiArgs struct {
	DecodeType interface{} // the point of Decode Type
}

// 反射版本的创建数组、append
func MongoQueryMulti(args MongoQueryMultiArgs) (rst interface{}, err error) {
	defer CheckPanic(&err)
	outType := reflect.TypeOf(args.DecodeType)
	outs := reflect.MakeSlice(reflect.SliceOf(outType), 0, 0)
	vals := reflect.New(outType.Elem()).Interface() // new 一个新对象
	outs = reflect.Append(outs, reflect.ValueOf(vals))
	rst = outs.Interface()
	return rst, nil
}

// UT 把panic转换成error
func CheckPanic(err *error) {
	p := recover()
	if p != nil {
		*err = fmt.Errorf("Catch Panic [%v]", p)
	}
}
