package Array

// 合并两个有序数组
// 输入:
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3

// 输出: [1,2,2,3,5,6]

func merge2Slice(nums1 []int, m int, nums2 []int, n int) {
	cur := m + n - 1
	i1 := m - 1
	i2 := n - 1
	for i1 >= 0 && i2 >= 0 { // 保证两个非负
		if nums1[i1] > nums2[i2] {
			nums1[cur] = nums1[i1]
			i1--
		} else {
			nums1[cur] = nums2[i2]
			i2--
		}
		cur--
	}
	if i2 >= 0 { // nums2这个数组有可能还有数,全部搬过来
		for i := 0; i <= i2; i++ {
			nums1[i] = nums2[i]
		}
	}
}

// 方法2，代码比较精巧
func merge2Slice10(nums1 []int, m int, nums2 []int, n int) {
	cur := m + n - 1
	i1 := m - 1
	i2 := n - 1
	for cur != i1 { // 核心，当i2插完时，i2 = -1,此时cur = i1,后续就不用插入了
		if i1 >= 0 && nums1[i1] > nums2[i2] {
			nums1[cur] = nums1[i1]
			i1--
		} else {
			nums1[cur] = nums2[i2] // num2插入
			i2--                   // 不用管i2为负的情况，当i2为负时，循环跳出了
		}
		cur--
	}
}
