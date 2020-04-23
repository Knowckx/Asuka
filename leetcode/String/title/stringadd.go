package title

import "strconv"

// 1. 定义⼤数:数字的⻓度超出了计算机int64的存储范围，需要使⽤字符串存储；计算两个⼤数相加，返回结果字符串。(代码实现，不限语⾔)

// 示例：
// 输⼊: ’3789742848912394213’ + ‘667213893899923292009992329’
// 输出: ‘667213897689666140922386542’

func addstring(s1 []byte, s2 []byte) {

	var long []byte
	var short []byte

	len1 := len(s1)
	len2 := len(s2)
	min = len1 // 确定最小长度的
	short, long = s1, s2
	if min > len2 {
		min = len2
		short, long = s2, s1
	}
	longlen := len(long)

	var out []byte
	fixInt := 0
	for i := 0; i < longlen; i++ {
		if i < min { //相同长度的数
			v1 := strconv.Atoi(short[i])
		} else {
			v1 = 0
		}
		v2 := strconv.Atoi(long[i])
		rst := v1 + v2 + fixInt
		if rst >= 10 {
			fixInt = 1
		} else {
			fixInt = 0
		}
		rsti = strconv.Itoa(rst)[1]
		out = append(out, rsti) // 该位和的结果
	}

	reverseString(out) // 反转
}

func reverseString(s []byte) {
	leng := len(s)
	for i, k := 0, leng-1; i < k; i, k = i+1, k-1 {
		s[i], s[k] = s[k], s[i]
	}
}
