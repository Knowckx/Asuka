package time

import (
	"fmt"
	"time"
)

// 输入一个时间，找到下一个周1的零时
func FindNextMonday(in time.Time) time.Time {
	d, m, s := in.Date()
	in = time.Date(d, m, s, 0, 0, 0, 0, time.Local) //当天的零时

	weekDay := in.Weekday()
	gapDay := (8 - weekDay) % 7
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
