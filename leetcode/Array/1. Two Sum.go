package Array

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:

// Given nums = [2, 7, 11, 15], target = 9,

// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

//nums := []int{3, 2, 4}

// 最佳解法 | 目前循环的数，无法和我map里的数组成target，那么就把这个循环的数放入map
func TwoSum1(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))
	outs := []int{}
	for i, v := range nums {
		v2 := target - v
		i2, ok := numsMap[v2]
		if !ok {
			numsMap[v] = i
			continue
		}
		outs = []int{i, i2}
	}
	return outs
}

//　暴力做法
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for i2 := i + 1; i2 < len(nums); i2++ {
			if nums[i]+nums[i2] == target {
				return []int{i, i2}
			}
		}
	}
	return nil
}
