package Array

import "fmt"

// Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.

// 判断数独矩阵是否有效
// 每行每列都必须不能有重复数字
// 每个3*3都不能有重复的数字

// Input:
// [
//   ["5","3",".",".","7",".",".",".","."],
//   ["6",".",".","1","9","5",".",".","."],
//   [".","9","8",".",".",".",".","6","."],
//   ["8",".",".",".","6",".",".",".","3"],
//   ["4",".",".","8",".","3",".",".","1"],
//   ["7",".",".",".","2",".",".",".","6"],
//   [".","6",".",".",".",".","2","8","."],
//   [".",".",".","4","1","9",".",".","5"],
//   [".",".",".",".","8",".",".","7","9"]
// ]
// Output: true

// Input:
// [
//   ["8","3",".",".","7",".",".",".","."],
//   ["6",".",".","1","9","5",".",".","."],
//   [".","9","8",".",".",".",".","6","."],
//   ["8",".",".",".","6",".",".",".","3"],
//   ["4",".",".","8",".","3",".",".","1"],
//   ["7",".",".",".","2",".",".",".","6"],
//   [".","6",".",".",".",".","2","8","."],
//   [".",".",".","4","1","9",".",".","5"],
//   [".",".",".",".","8",".",".","7","9"]
// ]
// Output: false
// Explanation: Same as Example 1, except with the 5 in the top left corner being
//     modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

// 	A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.
// The given board contain only digits 1-9 and the character '.'.
// The given board size is always 9x9.

func isValidSudoku(board [][]byte) bool {
	fmt.Println(board)
	return false
}

func isValidSudokuSlice(ins []byte) bool {
	bmap := make(map[byte]int, len(ins))
	for _, in := range ins {
		if in >= 0 && in <= 9 {
			bmap[in]++
			if bmap[in] > 1 {
				return false
			}
		}
	}
	return true
}

/*
使用3个二维数组，rows、cols、boxs分别用来存储行、列、子九宫格的情况
其中一维下标 n 对于3个二维数组分别表示：第 n 行，第 n 列，第 n 个子九宫格
其中二维下标 m 对于3个二维数组分别表示：在当前行、列、子九宫格的数字m
二维数组中的值则表示：该数字出现的次数（在本题中次数超过 1 次即代表重复）
举例：rows[2][5] = 1，第 2 行中数字 5 出现了 1 次
*/
