package problem19

import "fmt"

// TreeNode 树的结构体
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MirrorOfBinaryTree 二叉树的镜像
func MirrorOfBinaryTree(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	MirrorOfBinaryTree(root.Left)
	MirrorOfBinaryTree(root.Right)
}

func Print(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	Print(root.Left)
	Print(root.Right)
}
