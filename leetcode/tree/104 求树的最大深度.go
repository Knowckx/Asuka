package tree

// 求树的最大深度

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0 // root 可以是nil
	}
	curNode := []*TreeNode{root} // 数组A 表示本次遍历的结点集合
	nextNode := []*TreeNode{}    // 数组B 表示下一层的结点集合
	deep := 1
	for {
		for _, no := range curNode {
			if no.Left != nil {
				nextNode = append(nextNode, no.Left)
			}
			if no.Right != nil {
				nextNode = append(nextNode, no.Right)
			}
		}
		// 这一层的node遍历完了
		if len(nextNode) == 0 { // 如果下一层为0 那就跳出
			break
		}
		deep++
		curNode = nextNode
		nextNode = []*TreeNode{}
	}
	return deep
}
