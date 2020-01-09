package problem25

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// FindPath 寻找和为指定数的二叉树路径
func FindPath(root *TreeNode, sum int) [][]int {
	// 所有路径
	var res [][]int
	// 用来存储节点值的栈
	var solution []int
	var dfs func(root *TreeNode, sum int)

	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		rest := sum - root.Val
		if rest == 0 && root.Left == nil && root.Right == nil {
			solution = append(solution, root.Val)
			temp := make([]int, len(solution))
			copy(temp, solution)
			res = append(res, temp)
			solution = solution[:len(solution)-1]
			return
		}
		solution = append(solution, root.Val)
		dfs(root.Left, rest)
		dfs(root.Right, rest)
		solution = solution[:len(solution)-1]
	}
	dfs(root, sum)
	return res
}
