package Array

// Input:  [1,2,3,4,5,6,7] and k = 3
// Output: [5,6,7,1,2,3,4]
// 只能通过交换来做。

// 完美解0 反转
// [1,2,3,4,5,6,7] k=3
// [7,6,5,4,3,2,1] 1.全部反转一下 | 目的是为了把最后K个数的位置放到前边K个
// [5,6,7,4,3,2,1] 2.前K个数反转
// [5,6,7,1,2,3,4] 3.后len-K个数反转
func Rotate0(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
	return
}

// 工具函数，反转切片
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 简单1  暴力。K为几，那就插入几次
func Rotate1(nums []int, k int) {
	leng := len(nums)
	for ; k > 0; k-- {
		last := nums[leng-1]
		for i := 0; i < leng; i++ {
			temp := nums[i]
			nums[i] = last
			last = temp
		}
	}
	return
}

// 简单2  新开一个等长数组。K为几，就从这里切为两个切片，按次序放好。最后遍历一次，一个一个赋值就好。
// func Rotate2(nums []int, k int) {
// }
