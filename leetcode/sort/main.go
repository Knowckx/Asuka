package main

import "fmt"

func main() {
	bubbleSort()

}

// 冒泡排序
func bubbleSort() {
	ins := []int{10, 7, 1, 20, 40, 60, 2}
	for i := len(ins) - 1; i >= 0; i-- { // 外层循环,把所有数循环一遍 [0,n-1]
		for k := 0; k < i; k++ { // 内层循环 [0，i-1]
			if ins[k] < ins[k+1] {
				continue
			}
			ins[k], ins[k+1] = ins[k+1], ins[k]
		}
	}
	fmt.Println(ins)
}
