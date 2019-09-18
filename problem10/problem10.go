package problem10

// NumberOf1 统计1的个数
func NumberOf1(n int) int {
	count := 0
	flag := 1
	for flag != 0 {
		if flag&n != 0 {
			count++
		}
		flag = flag << 1
	}
	return count
}

// Ones 优化版NumberOf1
func Ones(n int) int {
	count := 0
	for n != 0 {
		count++
		n = n & (n - 1)
	}
	return count
}
