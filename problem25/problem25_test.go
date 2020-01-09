package problem25

import (
	"fmt"
	"testing"
)

func TestFindPath(t *testing.T) {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{3, nil, nil}, nil}}
	// 测试用例（有一条路径）
	res := FindPath(root, 8)
	fmt.Println(res)
	// 测试用例（有多条路径）
	res = FindPath(root, 7)
	fmt.Println(res)
	// 测试用例（没有满足路径）
	res = FindPath(root, 12)
	fmt.Println(res)
	// root为null二叉树（特殊）
	res = FindPath(nil, 12)
	fmt.Println(res)
}
