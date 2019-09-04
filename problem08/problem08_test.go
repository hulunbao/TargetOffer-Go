package problem08

import (
	"testing"
)

func TestMinNumberInRotateArray(t *testing.T) {
	arr := []int{4, 5, 6, 1, 2, 3}
	if minNumberInrotateArray(arr) != 1 {
		t.Errorf("expected answer is %v,but is %v", 1, minNumberInrotateArray(arr))
	}

	arr = []int{1, 2, 3, 4, 5, 6, 7}
	if minNumberInrotateArray(arr) != 1 {
		t.Errorf("expected answer is %v,but is %v", 1, minNumberInrotateArray(arr))
	}

	arr = []int{1, 0, 1, 1, 1}
	if minNumberInrotateArray(arr) != 0 {
		t.Errorf("expected answer is %v,but is %v", 0, minNumberInrotateArray(arr))
	}

	arr = []int{1, 1, 1, 0, 1}
	if minNumberInrotateArray(arr) != 0 {
		t.Errorf("expected answer is %v,but is %v", 0, minNumberInrotateArray(arr))
	}
}
