package problem15

import (
	"testing"
)

func TestFindKthToTail(t *testing.T) {
	l3 := &ListNode{3, nil}
	l2 := &ListNode{2, l3}
	l1 := &ListNode{1, l2}

	node := FindKthToTail(l1, 1)
	if node != l3 {
		t.Errorf("expected answer is %v,but is %v", l3, node)
	}

	node = FindKthToTail(l1, 2)
	if node != l2 {
		t.Errorf("expected answer is %v,but is %v", l2, node)
	}

	node = FindKthToTail(l1, 3)
	if node != l1 {
		t.Errorf("expected answer is %v,but is %v", l1, node)
	}

	node = FindKthToTail(l1, 0)
	if node != nil {
		t.Errorf("expected answer is %v,but is %v", nil, node)
	}

	node = FindKthToTail(l1, -1)
	if node != nil {
		t.Errorf("expected answer is %v,but is %v", nil, node)
	}

	node = FindKthToTail(nil, 2)
	if node != nil {
		t.Errorf("expected answer is %v,but is %v", nil, node)
	}

}
