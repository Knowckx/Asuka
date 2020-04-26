package title

/*
	21. 合并两个有序链表
	将两个有序链表合并
	输入：1->2->4, 1->3->4
	输出：1->1->2->3->4->4
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	out := &ListNode{} // 新建一个空引用，存放结果
	pre := out
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val { // 每次只对比l1和l2头部的两个节点，把小的那个插入到out中。
			pre.Next = l1
			l1 = l1.Next // l1头部被插入了。l1向后一步。
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	// 遍历完成后，可能有一个链表不为空，则直接加在后面（因为是有序的）
	if l1 != nil {
		pre.Next = l1
	}
	if l2 != nil {
		pre.Next = l2
	}
	return out.Next
}
