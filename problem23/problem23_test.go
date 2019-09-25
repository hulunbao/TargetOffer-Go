package problem23

import (
	"fmt"
	"testing"
)

func TestPrintFromTopToBottom(t *testing.T) {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	Print(root)
	fmt.Println(PrintFromTopToBottom(root))
}
