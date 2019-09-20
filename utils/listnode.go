package utils

import "fmt"

// ListNode 链表结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

// MakeListNode 构造单链表
func MakeListNode(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}

	head := &ListNode{Val: s[0]}
	index := head
	for i := 1; i < len(s); i++ {
		index.Next = &ListNode{Val: s[i]}
		index = index.Next
	}
	return head
}

// PrintListNode 打印链表
func PrintListNode(head *ListNode) {
	if head == nil {
		return
	}
	index := head
	for index != nil {
		fmt.Printf("%v->", index.Val)
		index = index.Next
	}
	fmt.Println()
}
