package Array

// 把0移动右边 | 只准操作原数组

// Input: [0,1,0,3,12]
// Output: [1,3,12,0,0]

// 遍历。遇到实数就按序号1，2，3，4往前扔。最后在k的位置后全补0
func moveZeroes(nums []int) {
	idx := 0
	for _, num := range nums {
		if num != 0 {
			nums[idx] = num
			idx++
		}
	}
	for i := idx; i < len(nums); i++ {
		nums[i] = 0
	}
}
