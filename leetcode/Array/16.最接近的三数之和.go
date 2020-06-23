package Array

import (
	"fmt"
	"math"
	"sort"
)

// S=[-1, 2, 1, -4]and target =1.
// 最接近 1 的三元组是 -1 + 2 + 1 = 2.

func Add3InPut(ins []int, tar int) []int {
	ins = []int{-1, 2, 1, -4}
	tar = 1
	sort.Ints(ins)

	// init
	res := ins[0:3]
	gap := math.Abs(float64(res[0] + res[1] + res[2] - tar))
	for le := 0; le < len(ins)-2; le++ {
		ri := len(ins) - 1
		mid := le + 1
		for mid < ri {
			ngap := ins[le] + ins[ri] + ins[mid] - tar
			if math.Abs(float64(ngap)) < gap {
				res = []int{ins[le], ins[mid], ins[ri]}
				mid++
				continue
			}
			if ngap > 0 {
				ri--
			} else {
				mid++
			}
		}
	}

	fmt.Println(res)
	return ins
}

// [-4 -1 1 2]

/*
	定义需要的东西  res = 目前提交的三个数  gap = 目前出现的最小差值
	时间复杂度：
		三数和，暴力遍历为n的三次方。
		目标是优化到N的二次方
	夹逼方法对应的时间复杂度为O(n),因此，左侧的left这个标识，需要向右移动
*/
