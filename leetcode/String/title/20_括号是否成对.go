package title

/*
	判断括号是否成对  腾讯50
	Input: "()[]{}"   	Output: true
	Input: "(]"			Output: false

*/

func IsValidPare(s string) bool {
	stk := &Stack{}                                        // 使用栈来遍历字符
	pareMap := map[rune]rune{'{': '}', '[': ']', '(': ')'} // 用map存放对应的字符配对
	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			stk.Push(c)
			continue
		}
		// c是一个右侧括号
		if stk.Count() == 0 {
			return false
		}

		peek := stk.Pop().(rune)
		if pareMap[peek] == c {
			continue
		}
		return false
	}
	return stk.Count() == 0
}
