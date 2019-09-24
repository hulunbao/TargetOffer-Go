package problem21

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	minstack := Constructor()
	minstack.Push(2)
	minstack.Push(5)
	minstack.Push(3)
	minstack.Push(6)
	minstack.Push(1)
	minstack.Pop()
	if minstack.GetMin() != 2 {
		t.Errorf("expected answer is %v,but is %v", 2, minstack)
	}
}
