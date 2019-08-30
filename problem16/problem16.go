package main

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
	p3 := &NodeList{}

	for p2 != nil {
		p3 = p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}
	head.Next = nil

	return p1
}

func print(head *NodeList) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
}

func main() {
	//test
	l3 := &NodeList{3, nil}
	l2 := &NodeList{2, l3}
	l1 := &NodeList{1, l2}
	fmt.Println("1 -> 2 -> 3")
	print(ReverseList1(l1))
	fmt.Println("\n")

	l2 = &NodeList{2, nil}
	l1 = &NodeList{1, l2}
	fmt.Println("1 -> 2 ->")
	print(ReverseList1(l1))
	fmt.Println("\n")

	l1 = &NodeList{1, nil}
	fmt.Println("1 -> ")
	print(ReverseList1(l1))
	fmt.Println("\n")

}
