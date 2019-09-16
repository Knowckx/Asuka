package time

import (
	"fmt"
	"testing"
)

func TestGetCusTimeGap(t *testing.T) {
	arg := CusTime{1, 13, 20, 0} // 定时任务在 周1,13:20:00

	var inTime Time = "2019-09-23 12:21:00"
	start := inTime.Time()
	// inTime := time.Now()
	rstGap := arg.NextGap(start)
	fmt.Println(rstGap)
	fmt.Println(start.Add(rstGap))
}
