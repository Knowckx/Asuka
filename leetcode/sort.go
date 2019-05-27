package topics

import (
	"fmt"
)

//快排
func qSortBegin() {
	arr := []int{40, 67, 89, 80, 21, 45, 37, 32, 59}
	qSort(arr)
	fmt.Printf("最后结果：%v", arr)
}

//快排实现
func qSort(arr []int) {
	first := 0 //头尾标识
	last := len(arr) - 1
	key := 0
	var tempArr []int //每次循环后的结果
	for _, v := range arr {
		if v >= arr[0] {
			tempArr = append(tempArr, v)
			continue
		}
		if v < arr[0] { //小值放左边
			tempArr = append([]int{v}, tempArr...)
			key++
		}
	}
	copy(arr, tempArr) //一趟排序完成
	fmt.Printf("一趟排序%v\n", arr)

	if key-first > 1 {
		qSort(arr[0:key])
	}
	if last-key > 1 {
		qSort(arr[key+1 : last+1])
	}
}
