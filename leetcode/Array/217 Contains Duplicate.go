package Array

import "github.com/Knowckx/Asuka/asuka"

//数组里是否有重复数字
// Input: [1,2,3,1]
// Output: true

func ContainsDuplicate(nums []int) bool {
	intMap := make(map[int]bool, len(nums))
	asuka.Display(len(intMap))
	for _, num := range nums {
		_, ok := intMap[num]
		if ok {
			return true
		}
		intMap[num] = true
	}
	return false
}
