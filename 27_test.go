package main

import (
	"reflect"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	var cases = []struct {
		in  []int
		val int

		out []int
	}{
		{
			[]int{},
			0,
			[]int{},
		},
		{
			[]int{1, 2, 5, 5, 3, 2, 4, 4, 4, 2, 2, 1},
			2,
			[]int{1, 5, 5, 3, 4, 4, 4, 1},
		},
		{
			[]int{1, 1},
			1,
			[]int{},
		},
		{
			[]int{1, 2, 2, 1},
			1,
			[]int{2, 2},
		},
	}
	for i, c := range cases {
		newL := removeElements(c.in, c.val)
		c.in = c.in[:newL]
		if !reflect.DeepEqual(c.in, c.out) {
			t.Errorf("Test case #%d: expected %v, received %v", i, c.out, c.in)
		}
	}
}
