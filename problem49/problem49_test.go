package problem49

import (
	"testing"
)

type para struct {
	str string
}

type ans struct {
	res int
}

type question struct {
	p para
	a ans
}

func TestMyAtoi(t *testing.T) {
	qs := []question{
		question{
			p: para{"42"},
			a: ans{42},
		},
		question{
			p: para{"-42"},
			a: ans{-42},
		},
		question{
			p: para{"0"},
			a: ans{0},
		},
		question{
			p: para{"0000"},
			a: ans{0},
		},
		question{
			p: para{""},
			a: ans{0},
		},
		question{
			p: para{"fdsfs"},
			a: ans{0},
		},
		question{
			p: para{"      42"},
			a: ans{42},
		},
		question{
			p: para{"      42asdfsdf"},
			a: ans{42},
		},
		question{
			p: para{"      -42asdfsdf"},
			a: ans{-42},
		},
		question{
			p: para{"9223372036854775808"},
			a: ans{0},
		},
		question{
			p: para{"-9223372036854775808"},
			a: ans{0},
		},
	}
	for _, q := range qs {
		if MyAtoi(q.p.str) != q.a.res {
			t.Errorf("expected answer is %v,but is %v", q.a.res, MyAtoi(q.p.str))
		}
	}
}
