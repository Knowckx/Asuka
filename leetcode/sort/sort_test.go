package sort

import (
	"container/heap"
	"fmt"
	"testing"
)

func Test_heapsort(t *testing.T) {
	lis := []int{9, 8, 6, 7, 2, 3}
	he := &IntHeap{}
	*he = lis
	heap.Init(he)   // 需要手动init
	fmt.Println(he) // lis[0] 为堆顶

	heap.Push(he, 1) // 可以看一下官方对于新加一个元素的up操作
	fmt.Println(he)

}
