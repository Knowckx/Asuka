package dp

// 斐波那契数列 | 爬梯子
// 非常标准的DP问题 DP[100] = DP[99] + DP[98]

// 除了最近的两次结果 其他的都不用保留
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	r1 := 1 //r1和r2永远指向最右的两个结果
	r2 := 2
	for i := 3; i <= n; i++ {
		r2, r1 = r1+r2, r2
	}
	return r2
}

// 使用递归，问题是很多cb[n]会被重复计算
func CB1(n int) int {
	if n <= 2 {
		return n
	}
	return CB1(n-1) + CB1(n-2)
}

// 数组 把每一次的结果都记录一下
func CB2(n int) int {
	if n <= 2 {
		return n
	}
	lis := make([]int, n+1)
	lis[1] = 1
	lis[2] = 2
	for i := 3; i <= n; i++ {
		lis[i] = lis[i-1] + lis[i-2]
	}
	return lis[n]
}
