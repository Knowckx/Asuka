package main

import (
	"fmt"

	dp "github.com/Knowckx/Asuka/leetcode/dp"
)

func main() {
	// dp.CB2(5)
	rst := dp.CB22(3, 2)
	fmt.Println(rst)
}

func Test() {
	s := []int{2, 3, 5, 7, 11, 13}
	_ = s
	// fmt.Println(s[-3:5])
}
