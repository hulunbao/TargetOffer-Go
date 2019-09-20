package problem17

import "fmt"

// ListNode 链表结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeList 合并有序链表递归
func MergeList(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	var MergeHead *ListNode
	if head1.Val < head2.Val {
		MergeHead = head1
		MergeHead.Next = MergeList(head1.Next, head2)
	} else {
		MergeHead = head2
		MergeHead.Next = MergeList(head1, head2.Next)
	}
	return MergeHead
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
