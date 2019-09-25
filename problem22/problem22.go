package problem22

import (
	"github.com/hulunbao/TargetOffer-Go/utils"
)

// IsPopOrder 是否为出栈序列
func IsPopOrder(pPush, pPop []int) bool {
	bPossible := false
	var stack utils.Stack
	if pPush != nil && pPop != nil && len(pPush) == len(pPop) {
		pushIndex := 0
		popIndex := 0
		for popIndex < len(pPop) {
			for stack.IsEmpty() || !Equels(stack, pPop[popIndex]) {
				if pushIndex == len(pPush) {
					break
				}

				stack.Push(pPush[pushIndex])
				pushIndex++
			}
			if !Equels(stack, pPop[popIndex]) {
				break
			}
			stack.Pop()
			popIndex++
		}
		if stack.IsEmpty() && popIndex == len(pPop) {
			bPossible = true
		}
	}
	return bPossible
}

// Equels 判断跟栈顶元素是否相等
func Equels(stack utils.Stack, value int) bool {
	top, _ := stack.Top()
	if top == value {
		return true
	}
	return false
}
