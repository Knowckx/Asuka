package Array

/*
	输入: [1,2,3]
	输出:
	[
	[1,2,3],
	[1,3,2],
	[2,1,3],
	[2,3,1],
	[3,1,2],
	[3,2,1]
	]
*/

/*
	全排序有点DP的味道
	假如一个切片 ins []int 长度为10
	在ins中取出一个数V0，则新的Ins为长度9
	res[10] = loop(res[9] + V0)
		对res[9]的结果遍历，每个结果添加上loop前取出的那个V0)

	res[9] = 对res[8]的结果遍历，每个结果加上loop前取出的V1)
*/

func permute(nums []int) [][]int {
	var res [][]int
	if len(nums) == 1 {
		return append(res, nums)
	}

	for i, num := range nums {
		temp := IntsRemoveAt(nums, i)
		lis := permute(temp)
		for _, li := range lis {
			res = append(res, append(li, num))
		}

	}
	return res
}
