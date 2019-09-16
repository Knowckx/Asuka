package time

import (
	"fmt"
	"time"
)

const FormatLayout1 = "2006-01-02"          // 10位
const FormatLayout2 = "2006-01-02 15:04:05" // 19位

/* ----------------- 字符串时间 Start ----------------- */

type Time string

func (in Time) Time() time.Time {
	layout := FormatLayout1
	if len(in) == 19 {
		layout = FormatLayout2
	}
	out, err := time.ParseInLocation(layout, string(in), time.Local)
	if err != nil {
		fmt.Printf("time.ParseInLocation,error %s,String %s\n", err, in)
		out = time.Time{}
	}
	return out
}

/* ----------------- 字符串时间 End   ----------------- */

// 输入一个时间，找到下一个周1的零时
func FindNextMonday(in time.Time) time.Time {
	d, m, s := in.Date()
	in = time.Date(d, m, s, 0, 0, 0, 0, time.Local) //当天的零时

	weekDay := in.Weekday()
	gapDay := (8 - weekDay) % 7 //距离周一的天数 | 重要
	ss := fmt.Sprintf("%dh", 24*gapDay)
	dur, _ := time.ParseDuration(ss)
	out := in.Add(dur)
	return out
}

// 时间 加以一个天数
func TimeAddDay(in time.Time, count int) time.Time {
	ss := fmt.Sprintf("%dh", 24*count)
	dur, _ := time.ParseDuration(ss)
	out := in.Add(dur)
	return out
}

/* ----------------- Time Job Start ----------------- */

const WEEKSEC int = 7 * 24 * 60 * 60 //一周的秒数

// 每周在此时间点执行
type CusTime struct {
	Day  int
	Hour int
	Min  int
	Sec  int
}

func NewCusTime(in time.Time) *CusTime {
	out := &CusTime{}
	out.Day = int(in.Weekday())
	out.Hour = in.Hour()
	out.Min = in.Minute()
	out.Sec = in.Second()
	return out
}

func (in *CusTime) ToSec() int {
	return in.Day*24*60*60 + in.Hour*60*60 + in.Min*60 + in.Sec
}

// 返回距离下一次执行点的Gap
func (in *CusTime) NextGap(start time.Time) time.Duration {
	now := NewCusTime(start).ToSec()
	tar := in.ToSec()
	sGap := (tar + WEEKSEC - now) % WEEKSEC
	ss := fmt.Sprintf("%ds", sGap)
	dur, _ := time.ParseDuration(ss)
	return dur
}

/* ----------------- Time Job End   ----------------- */
