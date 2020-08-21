package title

// 1. 定义⼤数:数字的⻓度超出了计算机int64的存储范围，需要使⽤字符串存储；计算两个⼤数相加，返回结果字符串。(代码实现，不限语⾔)

// 示例：
// 输⼊: ’3789742848912394213’ + ‘667213893899923292009992329’
// 输出: ‘667213897689666140922386542’

func addstring(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	maxLen := l1
	if l2 > maxLen {
		maxLen = l2
	}

	fix := 0
	rst := make([]byte, maxLen+1)

	for i := 0; ; i++ {
		if i >= l1 && i >= l2 {
			if fix != 0 {
				rst[0] = '1'
			} else {
				rst = rst[1:]
			}
			break
		}
		v1 := 0
		if i < l1 {
			v1 = int(num1[l1-i-1] - '0')
		}
		v2 := 0
		if i < l2 {
			v2 = int(num2[l2-i-1] - '0')
		}
		sum := v1 + v2 + fix
		newV := sum % 10
		if sum > 9 {
			fix = 1
		} else {
			fix = 0
		}
		rst[maxLen-i] = byte(newV) + '0'
	}
	return string(rst)
}
