package problem03

import (
	"testing"
)

type para struct {
	array [][]int
	value int
}

type ans struct {
	find bool
}

type question struct {
	p para
	a ans
}

func TestFind(t *testing.T) {
	question := []question{
		{
			p: para{
				array: [][]int{
					[]int{1,2,8,9},
					[]int{2,4,9,12},
					[]int{4,7,10,13},
					[]int{6,8,11,15},
				},
				value: 7,
			},
			a: ans{
				find:true,
			},
		},
		{
			p:para{
				array: [][]int{
					[]int{ 1, 2, 8, 9, },
					[]int{ 2, 4, 9, 12, },
					[]int{ 4, 6, 10, 13, },
					[]int{ 6, 8, 11, 15, },
				},
				value:7,
			},
			a:ans{
				find:false,
			},
		},

	}

	for _, q := range question{
		a, p := q.a, q.p
		if a.find != Find(p.array,p.value){
			t.Errorf("expected answer is %v,but is %v",a.find,Find(p.array,p.value))
		}

	}
}
