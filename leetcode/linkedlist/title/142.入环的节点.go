package title

import "fmt"

/*
	1.快慢指针向前走，只要是环，两个指针必然在环内相遇
	2.关键步骤，新指针重新在起点，指针2在相遇点，两者同时前进，再次相遇点就是环点
*/

func detectCycle(head *ListNode) *ListNode {
	fn := head
	sn := head
	for {
		if fn == nil || fn.Next == nil {
			return nil
		}
		fn = fn.Next.Next
		sn = sn.Next
		if fn == sn {
			fmt.Printf("P1:%d\n", fn.Val)
			break
		}
	}
	fn = head
	for {
		if fn == sn {
			fmt.Printf("P2:%d\n", fn.Val)
			return fn
		}
		fn = fn.Next
		sn = sn.Next
	}
}
