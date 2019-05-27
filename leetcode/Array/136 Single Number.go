package Array

// 找到数组里只出现一次的数字 | 额外条件，数组里重复的数最多出现2次 | 最多出现偶数次
// Input: [2,2,1]
// Output: 1

// 通用解法 第一遍map init, 第二遍遍历map找结果
func singleNumber(nums []int) int {
	temMap := make(map[int]int, len(nums))
	for _, num := range nums {
		v, ok := temMap[num]
		if !ok {
			temMap[num] = 1
			continue
		}
		temMap[num] = v + 1
	}

	out := 0
	for k, v := range temMap {
		if v == 1 {
			out = k
			break
		}
	}
	return out
}

// 含额外条件的最优解：位运算 异或（^）的自反性
func singleNumber2(nums []int) int {
	out := 0
	for _, num := range nums {
		out ^= num

	}
	return out
}
