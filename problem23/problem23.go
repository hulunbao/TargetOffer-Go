package problem23

import "fmt"

// TreeNode 二叉树结构体
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PrintFromTopToBottom 广度优先遍历
func PrintFromTopToBottom(root *TreeNode) []int {
	var queue []*TreeNode
	var res []int
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

	}
	return res
}

// Print 打印二叉树
func Print(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	Print(root.Left)
	Print(root.Right)
}
