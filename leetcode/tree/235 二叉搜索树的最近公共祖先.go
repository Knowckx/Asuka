package tree

// 二叉搜索树固有的性质，假如 左 根 右  是有序的！
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val < root.Val && q.Val > root.Val {
		return root
	}
	if p.Val < root.Val && q.Val < root.Val { // 这说明 p和q一定在左子树上
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
