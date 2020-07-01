package Array

import (
	"fmt"
	"testing"
)

func TestMerge2Slice(t *testing.T) {
	s1 := []int{1, 2, 3, 0, 0, 0}
	s2 := []int{2, 5, 6}
	fmt.Println("We Test")
	merge2Slice(s1, 3, s2, 3)
}

func TestThreeSum(t *testing.T) {
	s1 := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(s1)
	fmt.Println(res)
}

func Test_permute(t *testing.T) {
	s1 := []int{1, 2, 3}
	res := permute(s1)
	fmt.Println(res)
}

func Test_search33(t *testing.T) {
	s1 := []int{3, 1}
	res := search33(s1, 1)
	fmt.Println(res)
}

func Test_Add3InPut(t *testing.T) {
	nums := []int{1, 2, 5, 10, 11}
	tar := 12

	res := threeSumClosest(nums, tar)

	fmt.Println(res)
}

func Test_findKthLargest(t *testing.T) {
	nums := []int{1, 2, 5, 10, 11}
	tar := 1

	res := findKthLargest(nums, tar)

	fmt.Println(res)
}
