package linkedlist

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	s1 := makeLinkList([]int{1, 2, 3})
	s2 := makeLinkList([]int{1, 2, 10})
	ins := []*ListNode{s1, s2}

	n := mergeKLists(ins)
	fmt.Println(n)
}

func TestMakeLinkList(t *testing.T) {
	s1 := []int{1, 2, 3}
	n := makeLinkList(s1)
	fmt.Println(n)
}

func makeLinkList(lis []int) *ListNode {
	out := &ListNode{}
	tail := out
	for _, li := range lis {
		node := &ListNode{}
		node.Val = li
		tail.Next, tail = node, node
	}
	return out.Next
}
