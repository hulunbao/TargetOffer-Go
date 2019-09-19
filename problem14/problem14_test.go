package problem14

import (
	"fmt"
	"testing"
)

func TestReorderOddEven(t *testing.T) {
	num := []int{4, 2, 5, 7, 2, 1}
	ReorderOddEven(num)
	fmt.Println(num)

	num = []int{1, 3, 5, 2, 4, 6}
	ReorderOddEven(num)
	fmt.Println(num)

	num = []int{2, 4, 6, 1, 3, 5}
	ReorderOddEven(num)
	fmt.Println(num)

	num = []int{4}
	ReorderOddEven(num)
	fmt.Println(num)

	num = nil
	ReorderOddEven(num)
	fmt.Println(num)

	num = []int{}
	ReorderOddEven(num)
	fmt.Println(num)

}
