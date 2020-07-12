package digit

import (
	"sort"
)

/*
	请实现一个函数，输入一个正整数 n，返回与 n 组成数字完全相同，且小于 n 的最大整数。若不存在这样的数，返回 0。例如：
	●	165 => 156
	●	7123 => 3721
	●	8918 => 8891
	●	23 => 0
	●	103 => 0 (前导0不作为组成数字，因此31不符合要求
*/

// 89901 => 81990
func GetSecInt(input int) int {
	ins := make([]int, 0, 8)

	in := input
	for in != 0 {
		ri := in % 10
		ins = append(ins, ri) // 目前的数字排列的反向 [1,0,9,9,8]
		in = in / 10
	}

	sortIns := make([]int, len(ins)) // 排序结果 [0,1,8,9,9]
	copy(sortIns, ins)
	sort.Ints(sortIns)

	bigV := sortIns[len(sortIns)-1] // 9
	out := 0

	for i := len(ins) - 1; i >= 0; i-- {
		in = ins[i]
		if in != bigV {
			sortIns = RemoveOne(sortIns, in)
			out = out*10 + in
			continue
		}
		// 发现最大数，换成第二大的数
		if len(sortIns) == 1 {
			sortIns = []int{}
			out = 0
			break
		}
		secIdx := GetSecondBigIdx(sortIns)
		if secIdx == -1 {
			out = 0
			break
		}
		sec := sortIns[secIdx]
		sortIns = RemoveOne(sortIns, sec)
		out = out*10 + sec
		break
	}

	// 把剩余的数拼在后面 81 + 990
	for i := len(sortIns) - 1; i >= 0; i-- {
		v := sortIns[i]
		out = out*10 + v
	}

	return out
}

// 01899 --> 8
func GetSecondBigIdx(ins []int) int {
	idx := len(ins) - 2
	secIdx := -1
	big := ins[idx+1]
	for idx >= 0 {
		if ins[idx] != -1 && ins[idx] != big {
			secIdx = idx
			break
		}
		idx--
	}
	return secIdx
}

func RemoveOne(ins []int, tar int) []int {
	out := make([]int, 0, len(ins)-1)
	moved := false
	for _, in := range ins {
		if tar == in && !moved {
			moved = true
			continue
		}
		out = append(out, in)
	}
	return out
}
