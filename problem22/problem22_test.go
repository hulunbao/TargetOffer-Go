package problem22

import (
	"testing"
)

func TestIsOrder(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	out1 := []int{4, 5, 3, 2, 1}
	if !IsPopOrder(in, out1) {
		t.Error("error")
	}

	in = []int{1}
	out2 := []int{1}
	if !IsPopOrder(in, out2) {
		t.Error("error")
	}

	in = []int{1, 2, 3, 4, 5}
	out3 := []int{5, 4, 1, 2, 3}
	if IsPopOrder(in, out3) {
		t.Error("error")
	}

	in = []int{}
	out4 := []int{}
	if !IsPopOrder(in, out4) {
		t.Error("error")
	}

}
