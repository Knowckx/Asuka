package Array

import (
	"fmt"
	"sort"
)

// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

// 满足要求的三元组集合为：
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

/*
	本题解法：
	1.需要先排序，保持数字的大小关系
		[-1, 0, 1, 2, -1, -4] --> [-4, -1, -1, 0, 1, 2,]
		不要去重。这会导致(-1,-1,2) 这个结果的丢失
	2.先固定一个数，比如指针1固定在-4,外层循环每次前进一步
	3.固定一个数之后，再使用双指针夹逼，L=-1，R=2，向中间探索
	时间复杂度：排序的o(nlogn) + o(n`2)
*/

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums) // 排序
	fmt.Println(nums)
	tar := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > tar {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] { // 外层，重复数字不检查
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r { // 双指针
			if l > i+1 && nums[l-1] == nums[l] { // 重复数字跳过
				l++
				continue
			}
			if r < len(nums)-1 && nums[r+1] == nums[r] {
				r--
				continue
			}
			sum := nums[i] + nums[l] + nums[r]
			if sum == tar {
				re := []int{nums[i], nums[l], nums[r]}
				res = append(res, re)
				l++
				continue
			}
			if sum < tar {
				l++
			} else if sum > tar {
				r--
			}

		}

	}
	return res
}

/*
	套路：双指针收束
	前提：数组是有序的
	含义：有数组[1，2，3，4，5，6]
		l = 1 r = 6 根据l+r的结果，决定是l++还是r--
		总的时间复杂度为n，比暴力双指针遍历 -- [1，2] [1，3] [1，4]的时间复杂度 n*n要快很多。
	优势原因：
		暴力双指针遍历假如加入优化，也可以加快时间复杂度。
		比如要求l+r=5,事实上当L=2 R=3时，其实这里可以直接返回了。后面的都不需要遍历。
		但是这种优化，时间复杂度取决于于原始数据，并不好计算。
		双指针收束是一种时间复杂度更直观的优化算法，一定是O(n)
*/
