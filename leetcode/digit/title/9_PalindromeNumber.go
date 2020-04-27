package title

/*
	判断一个数字是不是回文的（对称）
	Input: 121
	Output: true

	Input: -121
	Output: false
*/

/*
	依赖 ti7,来实现对一半数的反转。
	奇数情况
		12721  当我们把最右的21反转成12时，此时是127 12
		127/10 = 12 就能发现这是回文的了。
	偶数情况
		127721
		127   127 直接相等
	负数情况
		-127 127- 直接pass
*/

func PalindInt(in int) bool {
	if in < 0 || (in%10 == 0 && in != 0) { // 1.负数不满足  2. 10 20 30这种也直接不满足
		return false
	}
	if in < 10 { // 个位数算回文
		return true
	}
	riPart := 0       // 存放右边部分的反转结果
	for in > riPart { // 核心：一半数反转时，已经可以判断结果了。  11211 当分解成 11 112算半折了
		ri := in % 10
		riPart = riPart*10 + ri
		in = in / 10 // 除10就相当于把最右边的数去掉了
	}
	// 出循环后，就是半边对折结果
	return riPart == in || riPart/10 == in
}
