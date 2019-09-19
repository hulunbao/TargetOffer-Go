package problem14

// ReorderOddEven 奇数放到偶数前面
func ReorderOddEven(num []int) {
	Reorder(num, isEven)
}

// Reorder 传递函数
func Reorder(num []int, fn func(num int) bool) {
	if num == nil || len(num) == 0 {
		return
	}
	left, right := 0, len(num)-1
	for left < right {
		for left < right && !fn(num[left]) {
			left++
		}
		for left < right && fn(num[right]) {
			right--
		}
		if left < right {
			num[left], num[right] = num[right], num[left]
		}
	}
}
func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}
