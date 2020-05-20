package title

/*
	字符串最长公共前缀
	输入: ["flower","flow","flight"]
	输出: "fl"

	输入：["",""]
*/

// 解法 [列]扫描 比较每个字符串的相同列是否相同

func LongestCommonPrefix(strs []string) string {
	res := ""
	if len(strs) == 0 {
		return res
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][0:i] // 直接通过idx切出来，无需保存中间结果
			}
		}
	}
	return strs[0]
}
