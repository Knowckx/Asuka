package linkedlist

/*
	20.04.17 思科电话面试 两个链表求和
	Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
	Output: 7 -> 8 -> 0 -> 7
*/

/*
	思路：
	求和运算是从最后一位的数字开始加，因此我们需要先反转一下单链表，然后再想加
*/

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1cur := Reverse(l1) // 先反转
	l2cur := Reverse(l2)
	var v1, v2 int
	fix := 0            // 进位
	newr := &ListNode{} // result
	cur := newr
	for {
		if l1cur == nil && l2cur == nil && fix == 0 {
			break
		}

		if l1cur == nil { // 取目前节点的两个值
			v1 = 0
		} else {
			v1 = l1cur.Val
			l1cur = l1cur.Next
		}

		if l2cur == nil {
			v2 = 0
		} else {
			v2 = l2cur.Val
			l2cur = l2cur.Next
		}

		sumv := v1 + v2 + fix // 求和
		v := 0
		if sumv >= 10 {
			fix = 1
			v = sumv - 10
		} else {
			fix = 0
			v = sumv
		}

		n := &ListNode{} // 结果保存成节点
		n.Val = v
		cur.Next = n
		cur = n

	}
	return Reverse(newr.Next)
}
