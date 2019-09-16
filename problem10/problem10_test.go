package problem10

import (
	"testing"
)

func TestNumberOf1(t *testing.T) {
	n := NumberOf1(0)
	if n != 0 {
		t.Errorf("expected answer is %v,but is %v", 0, n)
	}
	n = NumberOf1(-1)
	if n != 64 {
		t.Errorf("expected answer is %v,but is %v", 64, n)
	}
	n = NumberOf1(5)
	if n != 2 {
		t.Errorf("expected answer is %v,but is %v", 2, n)
	}
	n = NumberOf1(15)
	if n != 4 {
		t.Errorf("expected answer is %v,but is %v", 4, n)
	}

	if Ones(0) == 0 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}
	// int64
	if Ones(-1) == 64 {
		t.Log("Pass")
	} else {
		t.Error("Failed value: ", Ones(-1))
	}

	if Ones(8) == 1 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}

	if Ones(5) == 2 {
		t.Log("Pass")
	} else {
		t.Error("Failed")
	}

}
