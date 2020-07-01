package linkedlist

/*
	23. 合并多个有序链表
*/

func mergeKLists(lists []*ListNode) *ListNode {
	out := &ListNode{} // 新建一个空引用，存放结果
	tail := out

	for {
		minidx := -1
		for i := range lists {
			if lists[i] == nil {
				continue
			}
			if minidx == -1 {
				minidx = i
				continue
			}
			if lists[i].Val < lists[minidx].Val {
				minidx = i
			}
		}
		if minidx == -1 {
			break
		}
		tail.Next, tail = lists[minidx], lists[minidx] // 套路代码：关于增加尾节点，
		lists[minidx] = lists[minidx].Next
	}
	return out.Next
}
