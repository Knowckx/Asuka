package tree

/*
   -10
   / \
  9  20
    /  \
   15   7
   结果为42 也就是15+20+7
*/

var PathSumRes int // 记录结果

func maxPathSum(root *TreeNode) int {
	PathSumRes = root.Val
	GetMax(root)
	return PathSumRes
}

// 核心思想，递归。 本树的最大路径和 = 左树结果（非负） + 右树结果（非负） +  根结点
func GetMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := GetMax(root.Left)
	l = IntMax(l, 0) // 重点1：[左-2 中2 右2] 此时本树的结果应该是中加右=4.需要把左分支的负数去掉
	r := GetMax(root.Right)
	r = IntMax(r, 0)

	s := l + r + root.Val // 得到本树的结果
	if s > PathSumRes {
		PathSumRes = s
	}

	curmax := IntMax(l, r)
	sin := curmax + root.Val // 重点2：最后返回只能是半枝. 全树返回无法和上层的根结点构成单向笔划
	return sin
}

// 专门设立一个判断大小，美观代码
func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
