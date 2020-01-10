package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	1 -> 2 -> 3 -> 4 -> nil
	思维过程：
	抽取最小重复逻辑，假如目前在cur = 3位置
	我们希望cur.next = 2，那就需要一个pre变量记录2
	cur.next = 2会导致3.next丢失。需要一个temp先来记录3.next
	这些操作之后，3已经指向2了，我们需要前进
	pre,cur都要向前，cur正好 = temp pre=cur也就是目前的根结点
	下次循环cur == nil,正好退出

	下面的问题，loop前，pre和cur的初始状态多少为好。
	容易想到 pre = 1 cur = 2
	这样想也是好的，但是没有更极限
	假如cur = 1 pre = nil你试试呢？
	好像也行。
	对了，这里的loop从1位置就可以开始了。完美套入循环也没啥问题。
	完成。

*/

func Reverse1(head *ListNode) *ListNode {
	var pre *ListNode // 重要，ListNode类型的nil
	var cur = head
	for cur != nil {
		next := cur.Next //先保存
		cur.Next = pre
		pre, cur = cur, next
	}
	return pre
}

//     /*递归方式
//     在反转单链表的场景下，最好的处理方式就是从最后一个节点开始搞。
//     所以这里递归蛮合适的。
//     */
//     public ListListNode reverseList2(ListListNode head) {
//         if (head == null || head.next == null) { //初始次和递归次都需要检查
//             return head;
//         }
//         ListListNode p = reverseList2(head.next); //直接到倒数的结点
//         head.next.next = head; //4 → 5 改成了 4 ← 5
//         head.next = null; //没有这句，4 5就成了双向环
//         return p;
//     }
