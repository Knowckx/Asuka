package Array

// 找出数组中,出现次数大于 ⌊ n/2 ⌋ 的元素

func majorityElement(nums []int) int {
	cntMap := make(map[int]int, len(nums))
	for _, num := range nums {
		cntMap[num]++
		if cntMap[num] > len(nums)/2 {
			return num
		}
	}
	return 0
}
