package linkedlist

// 普通解 Map存放节点指针
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeMap := map[*ListNode]bool{} // 注意：需要存节点的地址 而不是Int
	for headA != nil {
		nodeMap[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		_, ok := nodeMap[headB]
		if ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// 最优解  len(A)+len(B) = len(B) + len(A)
// 所以把A的末尾接到B，把B的末尾接到A，再次遇到相同值就是交点
