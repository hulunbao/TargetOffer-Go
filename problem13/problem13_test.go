package problem13

import (
	"testing"
)

func TestDeletNode(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	head := MakeListNode(s)
	PrintListNode(head)
	head = DeleteNode(head, head.Next.Next)
	PrintListNode(head)

	s = []int{1, 2, 3, 4, 5, 6}
	head = MakeListNode(s)
	PrintListNode(head)
	head = DeleteNode(head, head)
	PrintListNode(head)

	s = []int{1, 2, 3, 4, 5, 6}
	head = MakeListNode(s)
	PrintListNode(head)
	head = DeleteNode(head, head.Next.Next.Next.Next.Next)
	PrintListNode(head)

	s = []int{1}
	head = MakeListNode(s)
	PrintListNode(head)
	head = DeleteNode(head, head)

	head = nil
	PrintListNode(head)
	head = DeleteNode(head, head)
	PrintListNode(head)

}
