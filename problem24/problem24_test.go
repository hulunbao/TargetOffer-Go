package problem24

import (
	"fmt"
	"testing"
)

func TestIsPostOrder(t *testing.T) {
	// 后序遍历
	post1 := []int{2, 9, 5, 16, 17, 15, 19, 18, 12}
	fmt.Println(post1, "is a postOrder ?  ", isPostOrder(post1))
	fmt.Printf("\n")
	// 没有右子树
	post2 := []int{2, 9, 5, 12}
	fmt.Println(post2, "is a postOrder ?  ", isPostOrder(post2))
	fmt.Printf("\n")
	// 没有左子树
	post3 := []int{13, 14, 15, 12}
	fmt.Println(post3, "is a postOrder ?  ", isPostOrder(post3))
	fmt.Printf("\n")
	// 只有一个节点
	post４ := []int{13}
	fmt.Println(post４, "is a postOrder ?  ", isPostOrder(post４))
	fmt.Printf("\n")

	post5 := []int{1, 16, 3, 13, 14, 15, 12}
	fmt.Println(post5, "is a postOrder ?  ", isPostOrder(post5))
	fmt.Printf("\n")
}
