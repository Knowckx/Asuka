package title

/*
	反转单链表
	1 -> 2 -> 3 -> 4 -> nil
	思维过程：
	抽取最小重复逻辑，假如目前在cur = 3位置
	我们希望cur.next = 2，那肯定需要一个pre变量记录2的位置
	cur.next = 2会导致3.next丢失。需要一个temp先来记录3.next
	这些操作之后，3已经指向2了，我们需要前进
	pre,cur都要向前，cur正好 = temp pre=cur也就是目前的根结点
	下次循环cur == nil,正好退出

*/

func Reverse(head *ListNode) *ListNode {
	var pre *ListNode // 重要，新建一个nil类型，可以在循环中让第一个节点指向nil
	var cur = head
	for cur != nil {
		next := cur.Next     //1.保存后一个元素
		cur.Next = pre       //2.后一个元素指向前一个元素
		pre, cur = cur, next //3.向前一步
	}
	return pre
}
