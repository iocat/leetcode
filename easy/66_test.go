package easy

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	cases := []struct {
		in  []int
		out []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{1, 2, 3, 4, 5, 7},
		},
		{
			[]int{9, 9, 9, 9, 9, 9, 9},
			[]int{1, 0, 0, 0, 0, 0, 0, 0},
		}, {
			[]int{1},
			[]int{2},
		},
		{
			[]int{1, 0, 9},
			[]int{1, 1, 0},
		},
		{
			[]int{9},
			[]int{1, 0},
		},
		{
			[]int{1, 1, 1, 1, 1},
			[]int{1, 1, 1, 1, 2},
		},
		{
			[]int{0, 0, 0, 0, 0, 0, 9},
			[]int{0, 0, 0, 0, 0, 1, 0},
		},
		{
			[]int{0, 9, 9, 9},
			[]int{1, 0, 0, 0},
		}, {
			[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, c := range cases {
		if res := plusOne(c.in); !reflect.DeepEqual(res, c.out) {
			t.Errorf("plusOne(%v) = %v, expected %v", c.in, res, c.out)
		}
	}
}
