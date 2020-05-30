package Array

// 数组的每个数字表示一个高，抽两个数字，与他们之间的距离，构成一个容器。求此容器的面积
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49

// 解法
// 使用套路：双指针收束
func maxArea(height []int) int {
	le := 0
	ri := len(height) - 1
	res := height[1] - height[0]
	for le < ri {
		min := GetMin(height[ri], height[le])
		area := min * (ri - le)
		if area > res {
			res = area
		}

		if height[le] < height[ri] {
			le++
		} else {
			ri--
		}
	}
	return res
}

func GetMin(v1 int, v2 int) int {
	if v1 <= v2 {
		return v1
	}
	return v2
}
