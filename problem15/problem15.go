package problem15

// ListNode 单链表结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

// FindKthToTail 找到链表倒数k个节点
func FindKthToTail(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return nil
	}
	Ahead := head

	for i := 0; i < k-1; i++ {
		if Ahead.Next != nil {
			Ahead = Ahead.Next
		} else {
			return nil
		}
	}

	Behind := head
	for Ahead.Next != nil {
		Ahead = Ahead.Next
		Behind = Behind.Next
	}
	return Behind

}
