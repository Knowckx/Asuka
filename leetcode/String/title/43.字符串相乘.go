package title

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

// 重点1 Char --> int   int('9' - '0') = 9
// 重点2 这道题的结果有可能大于int的范围，所以设置[]int数组,int[0]个位，int[1]表示十位
// 重点3 不用考虑连环进位，十位是11其实没有关系。下次加法会自己刷新叠上去。

/*
			99
	X		99
	--------------
			81  //  		0081
		   81   //  +810    0891
		   81   //  +810    1701
		  81    // +810		9801
	--------------
		  9801
*/

func Multiply3(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1, l2 := len(num1), len(num2)
	res := make([]int, l1+l2)
	for i := l1 - 1; i >= 0; i-- {
		v1 := int(num1[i] - '0')
		for k := l2 - 1; k >= 0; k-- {
			v2 := int(num2[k] - '0')
			sum := res[i+k+1] + v1*v2
			res[i+k+1] = sum % 10
			res[i+k] = sum/10 + res[i+k] // 最左是高位
		}
	}

	ss := ""
	for k, v := range res {
		if k == 0 && v == 0 {
			continue
		}
		//sk := strconv.Itoa(v)
		sk := string(v + '0') // 少导一个库
		ss += sk
	}

	return ss
}

// 边界分析，大家都在尾 l1-1 l2-1
// 此时res长为l1+l2,它的尾为 l1 + l2 - 1
