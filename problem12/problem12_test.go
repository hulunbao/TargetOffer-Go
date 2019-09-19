package problem12

import (
	"fmt"
	"testing"
)

func TestPrint1ToMaxNum(t *testing.T) {
	Print1ToMaxNum(3)
	fmt.Println()
	Print1ToMaxNum2(3)

	Print1ToMaxNum(0)
	fmt.Println()
	Print1ToMaxNum2(0)

	Print1ToMaxNum(-1)
	fmt.Println()
	Print1ToMaxNum2(-1)

}
