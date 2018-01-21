package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var cases = []struct {
		in  []int
		out []int
	}{
		{
			in:  []int{0},
			out: []int{0},
		},
		{
			in:  []int{0, 1, 2, 3},
			out: []int{0, 1, 2, 3},
		},
		{
			in:  []int{0, 0, 1, 1, 1, 2, 4, 5, 5, 5, 9},
			out: []int{0, 1, 2, 4, 5, 9},
		},
		{
			in:  []int{0, 0, 0},
			out: []int{0},
		},
		{
			in:  []int{1, 2, 2, 4, 5, 5, 5},
			out: []int{1, 2, 4, 5},
		},
	}
	for _, c := range cases {
		temp := make([]int, len(c.in))
		copy(temp, c.in)
		newL := removeDuplicates(temp)
		temp = temp[:newL]
		if !reflect.DeepEqual(temp, c.out) {
			t.Errorf("Error: removeDuplicates(%v) = expected %v, received %v", c.in, c.out, temp)
		}
	}

}
