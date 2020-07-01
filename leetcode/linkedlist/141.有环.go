package linkedlist

/*
	1.快慢指针向前走，只要是环，两个指针必然在环内相遇
*/

func hasCycle(head *ListNode) bool {
	fn := head
	sn := head
	for {
		if fn == nil || fn.Next == nil {
			return false
		}
		fn = fn.Next.Next
		sn = sn.Next
		if fn == sn {
			return true
		}
	}
}
