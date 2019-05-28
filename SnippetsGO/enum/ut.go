package enum

import (
	"fmt"
	"time"
)

const (
	FormatDayLayout       = "20060102"
	FormatDayLayoutDetail = "2006-01-02 15:04:05"
)

//S4 时间相关
var (
	S4TestMonth1Start Time = "2019-01-01 05:00:00"
	S4TestMonth3End   Time = "2019-03-30 05:00:00"
)

type Time string

func (in Time) Time() time.Time {
	out, err := time.ParseInLocation(FormatDayLayoutDetail, string(in), time.Local)
	if err != nil {
		fmt.Printf("time.ParseInLocation error %s args %s", err, in)
		out = time.Time{}
	}
	return out
}
