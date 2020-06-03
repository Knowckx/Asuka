package Array

/*
	输入一个旋转的排序数组 比如把[0,1,2,4,5,6,7,8,9] 变为 [4,5,6,7,8,9,0,1,2]
	搜索一个给定的目标值是否存在，不存在返回-1,存在返回下标
	要求算法时间复杂度必须是 O(log n)
*/

/*
	1.看到O(log n)，又是排序。想到二分查询.
	2.旋转排序数组的一个特点 至少有一半是有序的
	3.那我们的目标就是优先对有序的一半进行检查，只要检查两端就能知道tar是否在这半边中
	4.假如在半边中，很好，就是二分查询。
	5.重点来了，假如不是在半边中，那也很好啊，直接排除了一半数诶。也就完成了两分查询的效率来源。
	6.总结，所以，二分查询只要能保持局部有序就好了。
*/

// 90-99  0-40 | 40-90

func search33(nums []int, target int) int {
	le, ri := 0, len(nums)-1
	for le <= ri {
		mid := (le + ri) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[le] <= nums[mid] { // 说明左半段是有序的 | 需要=号，因为只有两个数的时候，le = 0 mid = 0
			if target >= nums[le] && target < nums[mid] {
				ri = mid - 1
			} else {
				le = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[ri] {
				le = mid + 1
			} else {
				ri = mid - 1

			}
		}
	}
	return -1
}
