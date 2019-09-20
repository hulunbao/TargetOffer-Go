package problem16

import "fmt"

// NodeList 链表结构体
type NodeList struct {
	Val  int
	Next *NodeList
}

// ReverseList 反转链表
func ReverseList(head *NodeList) *NodeList {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *NodeList
	cur := head

	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}

// ReverseList1 反转链表
func ReverseList1(head *NodeList) *NodeList {
	if head == nil {
		return head
	}

	p1 := head
	p2 := head.Next
	var p3 *NodeList

	for p2 != nil {
		p3 = p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}
	head.Next = nil

	return p1
}

// ReverseList2 反转链表
func ReverseList2(head *NodeList) *NodeList {
	var reverseHead *NodeList
	node := head
	var prev *NodeList
	for node != nil {
		next := node.Next
		if next == nil {
			reverseHead = node
		}
		node.Next = prev
		prev = node
		node = next
	}
	return reverseHead
}

func print(head *NodeList) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
}
