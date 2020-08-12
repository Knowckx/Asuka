
// 富途有一次
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
