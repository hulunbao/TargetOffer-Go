package problem17

import (
	"testing"
)

func TestMergeList(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 4, 6, 7}
	head1 := MakeListNode(s1)
	head2 := MakeListNode(s2)
	PrintListNode(head1)
	PrintListNode(head2)

	MergeHead := MergeList(head1, head2)
	PrintListNode(MergeHead)

	s1 = []int{1}
	s2 = []int{4}
	head1 = MakeListNode(s1)
	head2 = MakeListNode(s2)
	PrintListNode(head1)
	PrintListNode(head2)

	MergeHead = MergeList(head1, head2)
	PrintListNode(MergeHead)

	head1 = MakeListNode(nil)
	head2 = MakeListNode(nil)
	PrintListNode(head1)
	PrintListNode(head2)

	MergeHead = MergeList(head1, head2)
	PrintListNode(MergeHead)
}
