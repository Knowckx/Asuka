package time

import (
	"fmt"
	"testing"
	"time"

	"github.com/Knowckx/asukatime"
)

func Test_GetDayRange(t *testing.T) {
	input := time.Now()
	fmt.Println(input)
	tr := asukatime.GetDayRange(input)
	fmt.Println(tr.Start)
	fmt.Println(tr.End)
}
