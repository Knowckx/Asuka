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
