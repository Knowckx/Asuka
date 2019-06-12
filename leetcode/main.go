package main

import (
	"fmt"

	"github.com/Knowckx/Asuka/leetcode/Array"
)

func main() {
	nums := []int{3, 2, 4}
	tar := 6
	res := Array.TwoSum(nums, tar)
	fmt.Println(res)
}

func Test() {
	s := []int{2, 3, 5, 7, 11, 13}
	_ = s
	// fmt.Println(s[-3:5])
}
