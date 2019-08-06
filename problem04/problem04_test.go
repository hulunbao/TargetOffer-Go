package problem04

import (
	"testing"
)

type para struct {
	str    []byte
	length int
}

type ans struct {
	res []byte
}

type question struct {
	p para
	a ans
}

func TestReplaceSpace(t *testing.T) {
	qs := []question{
		question{
			p: para{
				str:    []byte{'a', ' ', 'b', ' ', 'c', 'x', 'x', 'x', 'x'},
				length: 5,
			},
			a: ans{
				res: []byte{'a', '%', '2', '0', 'b', '%', '2', '0', 'c'},
			},
		},
	}
	for _, q := range qs {
		p, a := q.p, q.a
		ReplaceSpace(p.str, p.length)
		if string(p.str) != string(a.res) {
			t.Errorf("expected answer is %v but is %v", string(a.res), string(p.str))
		}
	}
}
