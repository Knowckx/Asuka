package base

// 自定义栈结构
type Stack struct {
	lis    []interface{}
	length int
}

// 压栈操作
func (s *Stack) Push(in interface{}) {
	s.lis = append(s.lis, in)
	s.length++
}

// 出栈操作
func (s *Stack) Pop() interface{} {
	if s.length <= 0 {
		return nil
	}
	t := s.lis[s.length-1]
	s.lis = s.lis[:s.length-1]
	s.length--
	return t
}

// 返回栈顶元素
func (s *Stack) Peek() interface{} {
	if s.length <= 0 {
		return nil
	}
	return s.lis[s.length-1]
}

// 返回当前栈元素个数
func (s *Stack) Count() int {
	return s.length
}

// 清空栈
func (s *Stack) Clear() {
	s.lis = make([]interface{}, 8)
	s.length = 0
}

// 栈是否为空
func (s *Stack) IsEmpty() bool {
	return s.length == 0
}
