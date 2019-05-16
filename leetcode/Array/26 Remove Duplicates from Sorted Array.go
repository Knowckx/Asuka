package Array

// 题目：从一个已排序的数组里，得到其中不重复数字的数组
// 在原数组里进行值的交互与修改，返回不重复部分的长度 - 后续的部分可以不理会。
// 例如 {0, 0, 1, 1, 1, 2, 2, 3, 3, 4} ->  返回 {0,1,2,3,4,X,X,X,X} & int 5

// 最佳解法
// 双指针：i总是记录最新的那个不重复的数字的位置，j从第二个数字开始遍历。
// 每当发现新的不重复数字，把他赋值给i+1的位置
func RemoveDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
