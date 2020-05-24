package title

import (
	"bytes"
	"fmt"
	"math/rand"
)

// 141 链表有环
// 142 入环的节点

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
	root := &ListNode{}
	pre := root // 用于在循环中保持上一个节点
	for i := 0; i < leng; i++ {
		node := &ListNode{}
		r := rand.Intn(9)
		node.Val = r
		pre.Next = node
		pre = node
	}
	return root.Next
}
