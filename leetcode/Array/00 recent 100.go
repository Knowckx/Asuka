package Array

// 题目：一个长度为100的队列，每次加进来两个数，求最近100个数的合

// 队列的数组实现
type Queue []int

func (in *Queue) Put(v int) {
	newQ := append(*in, v)
	in = &newQ
}

func (in *Queue) Pop() int {
	tail := (*in)[0]
	newQ := (*in)[1:len(*in)]
	in = &newQ
	return tail
}

func (in *Queue) Size() int {
	return len(*in)
}
