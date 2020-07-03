package title

/*
	反转数字
	示例 2:
		输入: -123
		输出: -321
	示例 3:
		输入: 120
		输出:  21
*/

/*
	核心
	1.通过取余来获得最后一位的数字
	2.通过除10来让这个数去掉最后位 121 / 10 = 12
*/
func reverseInt(in int) int {
	out := 0           // 存放结果
	minInt := -1 << 31 // 左移为乘法 也就是 -21446464...
	maxInt := 1 << 31
	for in != 0 {
		ri := in % 10     // 取出个位 121 % 10 = 1
		out = out*10 + ri // 核心：每取的一个数字，就是原结果*10+此值
		if out <= minInt || out >= maxInt {
			return 0 // 取的过程中有可能超过范围了
		}
		in = in / 10 // 除10就相当于把最右边的数去掉了
	}
	return out
}
