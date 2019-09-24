package problem20

import (
	"fmt"
	"testing"
)

func TestPrintMatrixClockwisely(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(PrintMatrixClockwisely(matrix))

	matrix = [][]int{
		{1, 4, 5, 7, 7},
		{18, 4, 5, 7, 8},
		{17, 4, 5, 7, 9},
		{16, 4, 5, 7, 10},
		{15, 14, 13, 12, 11},
	}
	fmt.Println(PrintMatrixClockwisely(matrix))
}
