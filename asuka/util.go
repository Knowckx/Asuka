package asuka

import "math"

/*
	处理float64尾数不精确的问题
	例如 实际是0.06 但是内存里为表示0.05999999999999
	例如 实际是0.06 但是内存里为表示0.06000000000001
*/
func FixedFloat(num float64, precision int) float64 {
	if precision > 7 { // 防止越界
		return num
	}
	pow := math.Pow(10, float64(precision))
	round := float64(math.Round(num * pow))
	return round / pow
}
