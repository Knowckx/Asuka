package Array

/*
	输入: [3,2,1,5,6,4] 和 k = 2
	输出: 5
*/

func findKthLargest(nums []int, k int) int {
	buildMaxHeap(nums)
	size := len(nums)
	// 边界 k = 1,应该一次循环都不执行.前面的建堆操作已经确定了堆顶
	for i := size - 1; i >= size-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(nums[:i], 0)
	}
	return nums[0]
}

// 建堆
func buildMaxHeap(data []int) {
	heapSize := len(data)
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(data, i)
	}
}

// 从节点I开始进行Sift Down操作  若有子节点更大，需要交换调整
func maxHeapify(data []int, i int) {
	heapSize := len(data)
	le, ri, tar := i*2+1, i*2+2, i // tar 记录三者哪个是极限值
	if le < heapSize && data[le] > data[tar] {
		tar = le
	}
	if ri < heapSize && data[ri] > data[tar] {
		tar = ri
	}
	if tar != i {
		data[i], data[tar] = data[tar], data[i]
		maxHeapify(data, tar)
	}
}
