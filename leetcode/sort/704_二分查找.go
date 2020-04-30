package main

/*
	Binary Search
	Input: nums = [-1,0,3,5,9,12], target = 9
	Output: 下标为4
	没有找到这个数，直接返回 -1
*/

func BinSearch(nums []int, target int) int {

	startIdx := 0 // 初始的左右边界
	endIdx := len(nums) - 1

	for endIdx >= startIdx { // 假如左 = 右，那就是数组只有一个数
		curIdx := int((endIdx + startIdx) / 2)
		if nums[curIdx] == target {
			return curIdx
		}
		if nums[curIdx] < target {
			startIdx = curIdx + 1
			continue
		}
		if nums[curIdx] > target {
			endIdx = curIdx - 1
			continue
		}
	}
	return -1
}
