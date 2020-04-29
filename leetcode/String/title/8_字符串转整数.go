package title

import (
	"strings"
)

/*
	把字符转化为数字  腾讯50

	需要头尾去空格
	第一个字符可以为'+''-' 或者是普通数字'1'
	遇到第一个非数字，就可以直接返回
	int32的范围 数值范围为 [−231,  231 − 1]
*/

// 注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换，即无法进行有效转换。

// 在任何情况下，若函数不能进行有效的转换时，请返回 0 。

const (
	int32min = -1 << 31 // 刷题技巧
	int32max = 1<<31 - 1
)

func CusAtoi(str string) int {
	str = strings.TrimSpace(str) //去头尾的空白
	strlen := len(str)
	if strlen == 0 { // 空字符串的情况
		return 0
	}

	// 确定第一个字符的情况
	flag := 1
	start := 0
	if str[0] == '-' {
		flag = -1
		start = 1
	} else if str[0] == '+' {
		flag = 1
		start = 1
	} else if !(str[0] >= '0' && str[0] <= '9') {
		return 0
	}

	sum := 0
	for i := start; i < strlen; i++ {
		if !(str[i] >= '0' && str[i] <= '9') {
			break
		}
		cur := int(str[i] - '0') // 核心技巧，'9' - '0' = 9
		sum = sum*10 + cur*flag  // 注意点：flag需要放这里，保证sum就是负数，方便后续的比较

		if sum >= int32max {
			return int32max
		}
		if sum <= int32min {
			return int32min
		}
	}
	return sum
}
