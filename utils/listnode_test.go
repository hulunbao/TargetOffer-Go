package utils

import "testing"

func TestMakeListNode(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	head := MakeListNode(s)
	PrintListNode(head)
}

func TestPrintListNode(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	head := MakeListNode(s)
	PrintListNode(head)
}
