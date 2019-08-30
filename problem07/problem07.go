package problem07

import (
	"errors"

	"github.com/hulunbao/TargetOffer-Go/utils"
)

// Queue 队列结构
type Queue struct {
	in  utils.Stack
	out utils.Stack
}

// IsEmpty 是否为空
func (q Queue) IsEmpty() bool {
	return q.in.IsEmpty() && q.out.IsEmpty()
}

// Push 队列里添加数据
func (q *Queue) Push(value interface{}) {
	q.in.Push(value)
}

// Pop 队列弹出
func (q *Queue) Pop() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}

	if !q.out.IsEmpty() {
		value, _ := q.out.Pop()
		return value, nil
	}
	for !q.in.IsEmpty() {
		value, _ := q.in.Pop()
		q.out.Push(value)
	}

	value, _ := q.out.Pop()
	return value, nil
}
