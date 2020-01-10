package dp

/*
	0 1 1
	1 2 4

	m = 3
	n = 2
*/

func CB22(m, n int) int {
	lis2 := make([][]int, m) // 3列
	for i := 0; i < m; i++ {
		lis2[i] = make([]int, n) // 2行
	}
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if y == 0 && x == 0 {
				lis2[x][y] = 1
				continue
			}
			if y == 0 {
				lis2[x][y] = lis2[x-1][y]
				continue
			}
			if x == 0 {
				lis2[x][y] = lis2[x][y-1]
				continue
			}
			lis2[x][y] = lis2[x-1][y] + lis2[x][y-1]
		}
	}
	return lis2[m-1][n-1]
}
