package SnippetsGO

import "fmt"

// 一个长队列，每次我们只能消化掉step的东西，因此要分步做
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
