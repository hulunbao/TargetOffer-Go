package problem13

import (
	"fmt"
)

// ListNode  链表结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

// DeleteNode 删除链表节点O(1)
func DeleteNode(head, node *ListNode) *ListNode {
	if head == nil || node == nil {
		return nil
	}
	// 要删除的节点不是尾节点
	if node.Next != nil {
		index := node.Next
		node.Val = index.Val
		node.Next = index.Next
	} else if head == node {
		head = nil
	} else { // 删除节点在链表最后
		index := head
		for index.Next != node {
			index = index.Next
		}
		index.Next = nil
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
		fmt.Printf(" %v", index.Val)
		index = index.Next
	}
	fmt.Println()
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
