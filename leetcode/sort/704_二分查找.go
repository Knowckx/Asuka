package sort

import (
	"fmt"
	"math/rand"
)

/*
	Binary Search
	Input: nums = [-1,0,3,5,9,12], target = 9
	Output: 下标为4
	没有找到这个数，直接返回 -1
*/

func BinSearch(nums []int, target int) int {

	startIdx := 0 // 初始的左右边界
	endIdx := len(nums) - 1

	for endIdx >= startIdx { // 假如左 = 右，那就是数组只有一个数
		curIdx := int((endIdx + startIdx) / 2)
		if nums[curIdx] == target {
			return curIdx
		}
		if nums[curIdx] < target {
			startIdx = curIdx + 1
			continue
		}
		if nums[curIdx] > target {
			endIdx = curIdx - 1
			continue
		}
	}
	return -1
}

// 抽N个数
func GenRandomNums(alllen int, k int) []int {
	nums := map[int]bool{}
	for i := 0; i < alllen; i++ {
		nums[i] = true
	}

	outs := make([]int, 0, k)

	for i := 0; i < k; i++ {
		idx := rand.Intn(k) // 生成随机数
		outs = append(outs, idx)
		delete(nums, idx)
	}
	fmt.Println(outs)
	return outs

}

// 扫雷
func GenMineMap(length int, width int, mineCnt int) {
	GenRandomNums(50, 10)
	return

	mmap := make([][]int, width)
	for i := 0; i < width; i++ {
		out := make([]int, length)
		mmap[i] = out
	}
	fmt.Println(mmap)

	mineIdxMap := map[int]interface{}{}
	for k := 0; k < mineCnt; {
		alllen := length * width
		mineIndex := rand.Intn(alllen) // 生成随机数
		_, ok := mineIdxMap[mineIndex]
		if ok {
			continue
		}
		le := mineIndex / length
		wi := mineIndex / width
		mmap[le][wi] = 1 // 标记为雷
		k++
	}
	fmt.Println(mmap)

	for k := 0; k < width; k++ {
		for i := 0; i < length; i++ {
			if mmap[i][k] == 1 {
				continue
			}
			sum := 0
			if k-1 >= 0 && mmap[i][k-1] == 1 { // 左
				sum++
			}
			if i-1 >= 0 && k-1 >= 0 && mmap[i-1][k-1] == 1 { // 上
				sum++
			}

			if k+1 <= length-1 && mmap[i][k+1] == 1 { // 右
				sum++
			}

			if i+1 <= width-1 && k+1 <= length-1 && mmap[i+1][k] == 1 { // 下
				sum++
			}
			mmap[i][k] = sum

		}

	}
	fmt.Println(mmap)

}
