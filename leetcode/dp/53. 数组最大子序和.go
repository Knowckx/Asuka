package dp

/*
	输入数组Lis: [-2,1,-3,4,-1,2,1,-5,4],
	输出: 6
	解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/

/*
	得到公式 f(i)= max{ f(i−1)+Lis[i],Lis[i] }
	某个子序的最后一个元素是在I位置时，此子序的最大和，要麽是f(i-1)的和 + 新来的Lis[i],要麽就是直接使用Lis[i]新起一段。
	遍历完所有的f(I)之后，
	最后总体的结果一定是在某个f(i)中
*/

func maxSubArray(nums []int) int {
	res := nums[0]
	leng := len(nums)
	pre := nums[0]
	for i := 1; i < leng; i++ {
		v1 := pre + nums[i]
		v2 := nums[i]
		sumi := GetMax(v1, v2)
		res = GetMax(res, sumi)
		pre = sumi
	}

	return res
}

func GetMax(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
