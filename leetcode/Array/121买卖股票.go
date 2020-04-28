package Array

// 一组数字，表示每天的价格，求最好的入场和出场点后的利润
// 输入: [7,1,5,3,6,4]
// 输出: 5
// 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//      注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。

// 最优解 只需要遍历一遍
func maxProfit(prices []int) int {
	min := prices[0] // 记录一下目前的最小值
	maxP := 0        // 记录一下现今产生过的最大利润
	leng := len(prices)
	for i := 0; i < leng; i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}

		cur := prices[i] - min
		if cur > maxP {
			maxP = cur
		}
	}
	return maxP
}

// 暴力方法，遍历两遍。寻找所有组合里差值最大的
func maxProfit0(prices []int) int {
	max := 0
	leng := len(prices)
	for i := 0; i < leng; i++ {
		for k := i + 1; k < leng; k++ {
			cur := prices[k] - prices[i]
			if cur > max {
				max = cur
			}
		}
	}
	return max
}
