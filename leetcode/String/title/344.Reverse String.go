package title

// 反转字符串 空间需要为O(1)
// 最佳实践写法 -- 首尾标识向中间靠拢的场景
func reverseString(s []byte) {
	leng := len(s)
	for i, k := 0, leng-1; i < k; i, k = i+1, k-1 {
		s[i], s[k] = s[k], s[i]
	}
}
