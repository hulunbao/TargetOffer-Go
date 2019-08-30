package utils

import (
	"errors"
)

// Stack 栈类型
type Stack []interface{}

// Len 返回栈长度
func (stack Stack) Len() int {
	return len(stack)
}

// IsEmpty 是否为空
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

// Cap 返回容量
func (stack Stack) Cap() int {
	return cap(stack)
}

// Push 入栈
func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

// Top 获取栈定元素
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Out of index,len is 0")
	}
	return stack[len(stack)-1], nil
}

// Pop 出栈
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(*stack) == 0 {
		return nil, errors.New("Out of index,len is 0")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}
