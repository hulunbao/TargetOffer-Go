package problem18

// TreeNode 树结构体
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// HasSubTree 是否存在子树
func HasSubTree(root1, root2 *TreeNode) bool {
	result := false
	if root1 != nil && root2 != nil {
		if root1.Val == root2.Val {
			result = DoseTree1HaveTree2(root1, root2)
		}
		if !result {
			result = HasSubTree(root1.Left, root2)
		}
		if !result {
			result = HasSubTree(root1.Right, root2)
		}
	}
	return result
}

// DoseTree1HaveTree2 判断树1的子节点是否与树2相等
func DoseTree1HaveTree2(root1, root2 *TreeNode) bool {
	if root2 == nil {
		return true
	}
	if root1 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return DoseTree1HaveTree2(root1.Left, root2.Left) && DoseTree1HaveTree2(root1.Right, root2.Right)
}
