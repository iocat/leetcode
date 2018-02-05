package medium

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	cases := []struct {
		mat [][]int
		res []int
	}{
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		}, {
			[][]int{
				[]int{2, 2},
				[]int{3, 3},
				[]int{4, 4},
			},
			[]int{2, 2, 3, 4, 4, 3},
		},
		{[][]int{
			[]int{1, 1, 1, 1, 1},
			[]int{2, 2, 2, 2, 2},
			[]int{3, 3, 3, 3, 3},
			[]int{4, 4, 4, 4, 4},
			[]int{5, 5, 5, 5, 5},
			[]int{6, 6, 6, 6, 6},
		}, []int{1, 1, 1, 1, 1, 2, 3, 4, 5, 6, 6, 6, 6, 6, 5, 4, 3, 2, 2, 2, 2, 3, 4, 5, 5, 5, 4, 3, 3, 4},
		}, {
			[][]int{},
			[]int{},
		},
		{
			[][]int{[]int{1, 1}},
			[]int{1, 1},
		},
	}
	for _, c := range cases {

		if res := spiralOrder(c.mat); !reflect.DeepEqual(res, c.res) {
			t.Errorf("spiralOrder(%v) = %v, expected %v", c.mat, res, c.res)
		}
	}
}
