package main

import (
	"fmt"
)

//快排
func qSortBegin() {
	arr := []int{40, 67, 89, 80, 21, 45, 37, 32, 59}
	QSortByOneTurn(arr, 0, len(arr)-1)
	fmt.Printf("最后结果：%v", arr)
}

// 快排一次
func QSortByOneTurn(ins []int, left int, right int) {
	if left >= right {
		return
	}
	cur := left // 对比坐标
	std := ins[cur]

	for i := left + 1; i <= right; i++ {
		if ins[i] >= std {
			continue
		}
		temp := ins[i]
		ins[i], ins[cur+1] = ins[cur+1], ins[cur]
		ins[cur] = temp
		cur++
		// 这里 cur(基准) cur+1(基准+1)  i（发现一个比基础小的） 三个数 为了把i这个数放到基准的左侧
		// cur的位置放i，cur自己向前挪动一步，原来cur+1位置的数反正一定比cur大，可以放到i的位置
	}
	QSortByOneTurn(ins, left, cur-1)
	QSortByOneTurn(ins, cur+1, right)
}
