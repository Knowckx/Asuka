package main

import "fmt"

func main() {
	v1 := 3882.40
	v2 := 3881.61

	newRateIn := RateInput{v1, v2}
	rst := newRateIn.GetRate()
	fmt.Printf("Rate %0.5f\n", rst)

}

type RateInput struct {
	v1 float64
	v2 float64
}

func (in *RateInput) GetRate() float64 {
	if in.v2 == 0 {
		return 0
	}
	if in.v1 < in.v2 {
		in.v1, in.v2 = in.v2, in.v1
	}
	rst := (in.v1 - in.v2) / in.v2 * 100
	return rst
}
