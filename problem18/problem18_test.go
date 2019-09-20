package problem18

import (
	"fmt"
	"testing"
)

func TestHasSubTree(t *testing.T) {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	sub1 := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	sub2 := &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}
	sub3 := &TreeNode{2, nil, nil}

	fmt.Println(HasSubTree(root, sub1), HasSubTree(root, sub2), HasSubTree(root, sub3))
	// 不是子树
	sub4 := &TreeNode{2, &TreeNode{3, nil, nil}, nil}
	fmt.Println(HasSubTree(root, sub4))

	sub5 := &TreeNode{1, nil, nil}
	fmt.Println(HasSubTree(root, sub5))
	root = nil
	sub4 = nil
	fmt.Println(HasSubTree(root, sub4))
}
