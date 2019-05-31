package Array

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:

// Given nums = [2, 7, 11, 15], target = 9,

// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

func test(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))
	for i, v := range nums {
		numsMap[v] = i
	}

	outs := []int{}
	for v1, i1 := range numsMap {
		v2 := target - v1
		i2, ok := numsMap[v2]
		if ok {
			outs = []int{i1, i2}
			break
		}
	}
	return outs
}
