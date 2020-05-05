package tree

/*
	思维：中根遍历后的结果就是一个有序的数组，取第K个数。
*/

var cnt int
var res1 int // 保存递归中的结果

func kthSmallest(root *TreeNode, k int) int {
	cnt = k
	DFS(root)
	return res1
}

func DFS(root *TreeNode) {
	if root.Left != nil {
		DFS(root.Left)
	}
	cnt--
	if cnt == 0 {
		res1 = root.Val
		return
	} else if cnt < 0 {
		return
	}

	if root.Right != nil {
		DFS(root.Right)
	}
	return
}
