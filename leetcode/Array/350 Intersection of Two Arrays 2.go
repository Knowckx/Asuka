package Array

// 寻找交集

// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [4,9]

// Note:
// 重复几次就在结果出现几次
// 不需要按顺序
// Follow up:

// What if the given array is already sorted? How would you optimize your algorithm?
// What if nums1's size is small compared to nums2's size? Which algorithm is better?
// What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?

// 先init一个单map 然后对比
func intersect1(nums1 []int, nums2 []int) []int {
	numMap1 := make(map[int]int, len(nums1))
	for _, num := range nums1 {
		numMap1[num] += 1 //这个写法很棒
		// v, ok := numMap1[num]
		// if !ok {
		// 	numMap1[num] = 1
		// 	continue
		// }
		// numMap1[num] = v + 1
	}
	outs := []int{}
	for _, num := range nums2 {
		v, ok := numMap1[num]
		if ok && v != 0 {
			outs = append(outs, num)
			// numMap1[num] = v - 1
			numMap1[num] -= 1
			continue
		}

	}
	return outs
}
