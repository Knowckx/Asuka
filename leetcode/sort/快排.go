package sort

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
