package tree

import "fmt"

//
var initFlag bool

func isValidBST2(root *TreeNode) bool {
	lastV := 0
	if root == nil {
		return true
	}
	return GetTNodeVal(root, &lastV)
}

func GetTNodeVal(root *TreeNode, lastV *int) bool {
	ok := true
	fmt.Println("loop:", root.Val)

	if root.Left != nil {
		ok = GetTNodeVal(root.Left, lastV)
		if !ok {
			return ok
		}
	}

	fmt.Println(root.Val)

	if initFlag {
		*lastV = root.Val
		initFlag = false
	} else {
		if root.Val <= *lastV {
			return false
		}
		*lastV = root.Val
	}

	if root.Right != nil {
		ok = GetTNodeVal(root.Right, lastV)
		if !ok {
			return ok
		}
	}
	return true
}

func isValidBST(root *TreeNode) bool {
	initFlag = true
	return midloop(root)
}
	2
1		3
new in 2
1
stash out 2
Get Right 3
new in 3
3
new in 3
3

// 通过迭代来完成中序遍历
func midloop(root *TreeNode) bool {
	nodes := []*TreeNode{} // 遍历时保留节点
	cur := root
	lastV := 0
	for cur != nil || len(nodes) != 0 {
		// 逻辑步骤1，找到这个节点，最左边的节点，并把其他节点入栈
		fmt.Printf("new in %d\n", cur.Val)
		for cur.Left != nil {
			nodes = append(nodes, cur) // 入栈
			cur = cur.Left
			continue
		}
		fmt.Println(cur.Val) // 确定了左点
		if initFlag {
			initFlag = false
			lastV = cur.Val
		} else {
			if cur.Val <= lastV {
				return false
			}
			lastV = cur.Val
		}
		cur = nil

		for len(nodes) != 0 { // 继续遍历入栈点

			// 出栈一个
			leng := len(nodes)
			cur = nodes[leng-1] // 观察这个元素
			nodes = nodes[0 : leng-1]

			fmt.Printf("stash out %d\n", cur.Val) // 出栈的节点一定是中序节点，可以直接打印

			if cur.Val <= lastV {
				return false
			}
			lastV = cur.Val

			// 右节点问题
			if cur.Right == nil { // 情况1,假如是空,那就直接下个栈点吧
				cur = nil
				continue
			}
			// 右节点不为空，执行外层循环
			cur = cur.Right
			fmt.Printf("Get Right %d\n", cur.Val) // 出栈的节点一定是中序节点，可以直接打印
			break
		}
	}
	return true
}
