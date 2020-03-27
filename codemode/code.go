

	// 范式，分步去做
	step := 500
	total := len(ins)
	for index := 0; index < total; {
		next := index + step
		if next > total {
			next = total
		}
		AddSubscribePriceToTraders(ins[index:next])
		index = next
	}

	// XX时间后执行
	t := time.NewTimer(gap)
	<-t.C


	// 字符串时间

type Time string

func (in Time) Time() time.Time {
	out, err := time.ParseInLocation(FormatDayLayoutDetail, string(in), time.Local)
	if err != nil {
		log.Errw("time.ParseInLocation", "error", err, "args", in)
		out = time.Time{}
	}
	return out
}