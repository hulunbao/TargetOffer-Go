package problem19

import (
	"fmt"
	"testing"
)

func TestMirrorOfBinaryTree(t *testing.T) {
	//test
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	Print(root)
	fmt.Printf("\n")
	MirrorOfBinaryTree(root)
	Print(root)
	fmt.Printf("\n")

	root = &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	Print(root)
	fmt.Printf("\n")
	MirrorOfBinaryTree(root)
	Print(root)
	fmt.Printf("\n")

	root = &TreeNode{1, nil, nil}
	Print(root)
	fmt.Printf("\n")
	MirrorOfBinaryTree(root)
	Print(root)
	fmt.Printf("\n")

	root = nil
	Print(root)
	fmt.Printf("\n")
	MirrorOfBinaryTree(root)
	Print(root)
	fmt.Printf("\n")
}
