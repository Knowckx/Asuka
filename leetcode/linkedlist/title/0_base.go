package title

import (
	"bytes"
	"fmt"
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (in *ListNode) Print() {
	outs := bytes.Buffer{}
	if in != nil {
		s := fmt.Sprintf("%d", in.Val)
		outs.WriteString(s)
		in = in.Next
	}
	for in != nil {
		s := fmt.Sprintf(" --> %d", in.Val)
		outs.WriteString(s)
		in = in.Next
	}
	fmt.Println(outs.String())
}

func GenListNode(leng int) *ListNode {
	pre := &ListNode{}
	for i := 0; i < leng; i++ {
		node := &ListNode{}
		r := rand.Intn(100)
		node.Val = r
		pre.Next = node
		pre = node
	}
	return pre.Next
}
