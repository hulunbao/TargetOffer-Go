package problem12

import "fmt"

// Print1ToMaxNum 打印一到最大的n位数
func Print1ToMaxNum(n int) {
	if n < 0 {
		return
	}
	number := make([]int, n)
	for !IncrementNumber(number) {
		PrintNumber(number)
	}
}

// IncrementNumber 模拟增加字符串判断是否为最大
func IncrementNumber(number []int) bool {
	isOverflow := false
	nTakeOver := 0
	nLength := len(number)
	for i := nLength - 1; i >= 0; i-- {
		nSum := number[i] + nTakeOver
		if i == nLength-1 {
			nSum++
		}

		if nSum >= 10 {
			if i == 0 {
				isOverflow = true
			} else {
				nSum -= 10
				nTakeOver = 1
				number[i] = nSum
			}
		} else {
			number[i] = nSum
			break
		}
	}
	return isOverflow
}

// PrintNumber 打印
func PrintNumber(number []int) {
	isBeginning := true
	for i := 0; i < len(number); i++ {
		if isBeginning && number[i] != 0 {
			isBeginning = false
		}
		if !isBeginning {
			fmt.Print(number[i])
		}
	}
	fmt.Print(" ")
}

// Print1ToMaxNum2 全排列计算
func Print1ToMaxNum2(n int) {
	if n < 0 {
		return
	}

	number := make([]int, n)
	for i := 0; i < 10; i++ {
		number[0] = i
		Print1ToMaxNumRecursively(number, n, 0)
	}

}

// Print1ToMaxNumRecursively 递归输出
func Print1ToMaxNumRecursively(number []int, length, index int) {
	if index == length-1 {
		PrintNumber(number)
		return
	}

	for i := 0; i < 10; i++ {
		number[index+1] = i
		Print1ToMaxNumRecursively(number, length, index+1)
	}
}
