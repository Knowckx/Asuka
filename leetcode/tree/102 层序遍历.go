package tree

// 广度优先 打印二叉树
func levelOrder(root *TreeNode) [][]int {
	outs := [][]int{}
	if root == nil {
		return outs
	}

	cusNodes := []*TreeNode{root}
	nextNodes := []*TreeNode{}
	for {
		out := []int{}
		for _, no := range cusNodes {
			out = append(out, no.Val)
			if no.Left != nil {
				nextNodes = append(nextNodes, no.Left)
			}
			if no.Right != nil {
				nextNodes = append(nextNodes, no.Right)
			}
		}
		outs = append(outs, out)
		if len(nextNodes) == 0 {
			break
		}
		cusNodes = nextNodes
		nextNodes = []*TreeNode{}

	}
	return outs

}
