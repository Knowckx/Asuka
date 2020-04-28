package Array

// 给定一串数字的数组，相当于股票每天的价格，以求得理论上的做多最大利润。

// Input: [7,1,5,3,6,4]
// B1 S5  B3 S6  总利润7

// 最佳解法 遍历，前后两个值的差，为正就是有效利润！这些有效利润的和就是总利润
func maxProfit2(prices []int) int {
	sum := 0
	for i := 0; i < len(prices)-1; i++ {
		pf := prices[i+1] - prices[i]
		if pf > 0 {
			sum += pf
		}
	}
	return sum
}
