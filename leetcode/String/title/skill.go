package title

import "fmt"

// 记录技巧
func SkillResult() {
	// byte '8' --> int 8 ??
	v1 := int('8' - '0')
	fmt.Println(v1)

	// int 8 --> byte '8' ??
	// 记住一个点 int 49 = byte '1'
	by0 := byte(8 + '0')
	fmt.Println(string(by0)) // 输出 "8"

	// []byte and string ??
	s1 := "1234" // s1[0] == bys[0] == '1' == ascii 49
	bys := []byte(s1)
	fmt.Println(len(s1), len(bys)) // 长度是相同的！其实两者完全等价
	fmt.Println(s1[0], bys[0])     // 结果是相同的，都是'1'，也就是ascii 49

}

// test入口
func DoSkill() {
	SkillResult()
}

// 一些基础的小技巧
func Skill1() {

}
