package title

/*
	给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
*/

/*
	4 ~ 100
	8 ~ 1000
	16 ~ 10000  // 最高位为1, 其余全是0
	15 ~ 01111  // 二进制下都是1;

	n&(n-1)必为0
*/

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}
