package problem49

import (
	"math"
)

// MyAtoi 手动实现Atoi函数
func MyAtoi(str string) int {
	res := 0
	flag := 1
	index := 0

	// 空字符串
	if len(str) == 0 {
		return res
	}
	// 前面有空格的情况
	for index < len(str) && str[index] == ' ' {
		index++
	}

	// + -的情况
	if index < len(str) && str[index] == '+' {
		index++
	} else if index < len(str) && str[index] == '-' {
		index++
		flag = -1
	}

	// 处理字符串
	for index < len(str) && str[index] <= '9' && str[index] >= '0' {
		if str[index] == '0' {
			res = 10*res + int(str[index]-'0')
			index++
		} else {
			res = 10*res + int(str[index]-'0')
			index++
		}
		if res*flag > math.MaxInt32 || res*flag < math.MinInt32 {
			return 0
		}

	}

	return res * flag
}
