package digit

import (
	"testing"
)

func TestGetSecInt(t *testing.T) {
	input := []int{
		165,
		7123,
		8918,
		23,
		103,
	}
	ans := []int{
		156,
		3721,
		8891,
		0,
		0,
	}
	for i := 0; i < len(input); i++ {
		an := GetSecInt(input[i])
		if an != ans[i] {
			t.Errorf("worne answer:%d => %d", input[i], ans[i])
		}
	}
}
