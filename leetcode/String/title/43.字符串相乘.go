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

// int[] res = new int[num1.length() + num2.length()]; // cap  4
// for (int i = num1.length() - 1; i >= 0; i--) { // i = 1
// 	int n1 = num1.charAt(i) - '0';
// 	for (int j = num2.length() - 1; j >= 0; j--) { // j = 2
// 		int n2 = num2.charAt(j) - '0';
// 		int sum = (res[i + j + 1] + n1 * n2);  sum =  8
// 		res[i + j + 1] = sum % 10;  8
// 		res[i + j] += sum / 10;   0
// 	}
// }

// 重点1 '9' - '0' = 9
// 重点2 因为这个结果数组很长，所以设置一个[]int数组,int[0]表示个位结果，int[1]表示十位结果
// 重点3 每次有个数加进来，不需要考虑进位的问题，因为本位就算暂时是11，也可以放得下。

/*
			99
	X		99
	--------------
			81  //  0081
		   81   //  +810
		   		//= 0891
		   81   //  +810
		   		//= 1701
		  81    // +810
		  		//= 9801
	--------------
		  9801
*/

func Multiply3(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1, l2 := len(num1), len(num2)
	ll := l1 + l2 - 2
	res := make([]int, l1+l2, 0)
	for i := l1 - 1; i >= 0; i-- {
		v1 := int(num1[i] - '0')
		for k := l2 - 1; k >= 0; k++ {
			v2 := int(num2[k] - '0')
			sum := res[ll-i-k] + v1*v2
			res[ll-i-k] = sum % 10
			res[ll-i-k+1] = sum/10 + res[ll-i-k+1]
		}
	}
	fmt.Println(res)
	return "0"
}
