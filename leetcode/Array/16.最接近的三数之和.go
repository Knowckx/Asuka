package Array

import (
	"sort"
)

// S=[-1, 2, 1, -4]and target =1.
// 最接近 1 的三元组是 -1 + 2 + 1 = 2.

// 1, 2, 5, 10, 11 | 12     预期 13
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	// init
	res := nums[0] + nums[1] + nums[2]
	for le := 0; le < len(nums)-2; le++ {
		ri := len(nums) - 1
		mid := le + 1
		for mid < ri {
			sum := nums[le] + nums[ri] + nums[mid]
			if distance(sum, target) < distance(res, target) {
				res = sum
			}
			if sum-target > 0 {
				ri--
			} else {
				mid++
			}
		}
	}
	return res
}

// 要比math.Abs(float(res-target))快速
func distance(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

/*
	定义需要的东西  res = 目前提交的三个数  gap = 目前出现的最小差值
	时间复杂度：
		三数和，暴力遍历为n的三次方。
		目标是优化到N的二次方
	夹逼方法对应的时间复杂度为O(n),因此，左侧的left这个标识，需要向右移动
*/
