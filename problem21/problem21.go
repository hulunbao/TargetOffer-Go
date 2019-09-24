package problem21

// MinStack 最小数队列
type MinStack struct {
	stack []item
}

type item struct {
	value, min int
}

// Constructor 构造最小队列
func Constructor() MinStack {
	return MinStack{}
}

// Push 入队
func (s *MinStack) Push(value int) {
	min := value
	if len(s.stack) > 0 && s.GetMin() < value {
		min = s.GetMin()
	}
	s.stack = append(s.stack, item{value: value, min: min})
}

// Pop 出栈
func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
}

// Top 返回栈顶元素
func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1].value
}

// GetMin 返回栈最小的元素
func (s *MinStack) GetMin() int {
	return s.stack[len(s.stack)-1].min
}
