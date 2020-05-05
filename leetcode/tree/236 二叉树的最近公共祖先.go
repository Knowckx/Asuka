package tree

var res *TreeNode

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// 思路是只要root,left,right，三个点发现一个存在命中，直接呈交上去
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	n1 := lowestCommonAncestor2(root.Left, p, q)
	n2 := lowestCommonAncestor2(root.Right, p, q)
	if n1 != nil && n2 != nil {
		return root
	} else if n1 != nil {
		return n1
	} else if n2 != nil {
		return n2
	}
	return nil
}

/*
	  5
	1	2
	假如5和2是目标数字，我原本以为需要同时对比root和左右三种情况，进行处理。
	实际上不需要。当检查root = 5时，直接就可以返回了。

	假如2是5的子树，那么直接返回5是最终答案

	假如2不是5的子树，返回的5节点，是上层root的一个子树解
*/
