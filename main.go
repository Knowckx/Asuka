package main

import (
	"fmt"
)

func main() {
	fmt.Println("asuka hello world")

	// merge2Sliceı˜
	// test()

}

func test() {
	tarList := make([]int, 11)
	fmt.Print(tarList)

	step := 3
	for index := 0; index < len(tarList); {
		next := index + step
		if next > len(tarList) {
			next = len(tarList)
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

type AA struct {
	v1 int
	v2 string
	li []int
}

func (A AA) Copy() AA {
	return A
}
