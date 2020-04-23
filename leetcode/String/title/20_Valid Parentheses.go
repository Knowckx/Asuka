package title

import "fmt"

/*
	判断括号是否成对
	Input: "()[]{}"   	Output: true
	Input: "(]"			Output: false

	从左向右遍历，把字符压到栈里
*/

func isValidPare(s string) bool {
	for _, c := range s {
		fmt.Println(c)
	}
	return false
}

// class Solution:
//     def isValid(self, s: str) -> bool:
//         dic = {'{': '}',  '[': ']', '(': ')', '?': '?'}
//         stack = ['?']
//         for c in s:
//             if c in dic: stack.append(c)
//             elif dic[stack.pop()] != c: return False
//         return len(stack) == 1
