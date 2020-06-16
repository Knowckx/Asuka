package title

import (
	"fmt"
	"math"
	"strconv"
)

/*
	输入: num1 = "2", num2 = "3"
	输出: "6"

	输入: num1 = "123", num2 = "456"
	输出: "56088"
*/

/*
			12
	X		34
	--------------
			08
		   04
		   06
		  03
	--------------
		   408
*/

func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	res := 0
	l1, l2 := len(num1), len(num2)
	for i := 0; i <= l1-1; i++ {
		for k := 0; k < l2; k++ {
			v1 := int(num1[l1-1-i] - '0')
			v2 := int(num2[l2-1-k] - '0')
			vv := v1 * v2 * int(math.Pow10(i+k))
			fmt.Println(vv)
			res += vv
		}
	}
	fmt.Println(res)
	return strconv.Itoa(res)
}

// ultiply1(num1 string, num2 string) string {
// 	if num1 == "0" || num2 == "0" {
// 		return "0"
// 	}
// 	l1, l2 := len(num1), len(num2)
// 	l := l1 + l2
// 	carry := make([]byte, l)
// 	for i := l1; i > 0; i-- {
// 		index := l2 + i - 1
// 		n1 := num1[i-1] - '0'
// 		for j := l2; j > 0; j-- {
// 			n := carry[index] + n1*(num2[j-1]-'0') //计算
// 			carry[index] = n % 10                  //当前值
// 			index--
// 			carry[index] += n / 10 //进位
// 		}
// 	}

// 	j := -1
// 	for i := 0; i < l; i++ {
// 		if carry[i] != 0 && j == -1 {
// 			j = i
// 		}
// 		carry[i] += '0'
// 	}
// 	return string(carry[j:])
