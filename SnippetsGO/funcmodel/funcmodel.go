package funcmodel

import "fmt"

// 把[1,2,3]转换成"1,2,3"
// 函数模式：字符串拼接，第一个数时需要特别处理
func IntsToString(bids []int) string {
	out := ""
	if len(bids) == 0 {
		return out
	}
	for i, bid := range bids {
		if i == 0 {
			out = fmt.Sprintf("%d", bid)
			continue
		}
		out = fmt.Sprintf("%s,%d", out, bid)
	}
	return out
}

// 分步做
func DoStep() {
	tarList := make([]int, 11)
	fmt.Print(tarList)

	step := 3
	total := len(tarList)
	for index := 0; index < total; {
		next := index + step
		if next > total {
			next = total
		}
		DoList(tarList[index:next])
		index = next
	}
	fmt.Print(tarList)
}

func DoList(ins []int) {
	for k := range ins {
		ins[k] = 5
	}
}
