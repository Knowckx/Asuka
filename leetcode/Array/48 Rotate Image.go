package Array

/*
	需要在原数组上进行修改
*/
// len = 3
// [
//   [1,2,3],
//   [4,5,6],
//   [7,8,9]
// ],

// rotate the input matrix in-place such that it becomes:
// [
//   [7,4,1],
//   [8,5,2],
//   [9,6,3]
// ]

// 分析 3*3 N=3
// 0,0 --> 0,2 --> 2,2 --> 2,0
// 0,1 --> 1,2 --> 2,1 --> 1,0
// 结论：[i,k] --> [k,N-i-1]

// len = 4
// [
//   [1,2,3,4],
//   [2,4,5,6],
//   [3,7,8,9],
//   [8,7,5,2],
// ],

func rotate(matrix [][]int) {
	N := len(matrix)
	temp := [][]int{} //复制
	for i := 0; i < N; i++ {
		s := []int{}
		for k := 0; k < N; k++ {
			s = append(s, matrix[i][k])
		}
		temp = append(temp, s)
	}
	for i := 0; i < N; i++ {
		for k := 0; k < N; k++ {
			matrix[k][N-i-1] = temp[i][k]
		}
	}
}

// 假如要求 空间复杂度为1，不建立额外矩阵 会比较难，需要扣。

// 遍历次数
// 外层遍历为 N/2次 N=2 1次 N=3 1次 N=4 两次
// 内层 从 0,0 到 0,N/2
// 每次换四个数
// 外层 N = 4 时 0，0 1，1
