package title

/*
	回文串有两种形式
	输入: "babad"  形态1 长度为奇数中间是一个值
	输出: "bab"

	输入: "cbbd"  形态2 长度为偶数，中间是两个值相同
	输出: "bb"
*/

// 中心向两边扩展
func LongestPalindrome(s string) string {
	if s == "" || len(s) == 0 {
		return ""
	}
	var le, ri int
	for i := 0; i < len(s)-1; i++ {
		len1 := CheckPaLen(s, i, i)
		len2 := CheckPaLen(s, i, i+1)
		if ri-le < len1 {
			ri = i + len1/2 // aaa r-l=2
			le = i - len1/2
		}
		if ri-le < len2 { // abba r-l=3
			ri = i + (len2-1)/2 + 1
			le = i - (len2-1)/2
		}
	}
	return s[le : ri+1]
}

// 从idx这个点向两边扩展，最长的长度 | 会有两种形态
func CheckPaLen(s string, l int, r int) int {
	var le, ri int
	for l >= 0 && r <= len(s)-1 {
		if s[l] != s[r] {
			break
		}
		le, ri = l, r // 确定成功一次，才可以跟踪一次
		l--
		r++
	}
	return ri - le // 返回r-l，因为上层函数记录的也是r和l,保持统一
}
