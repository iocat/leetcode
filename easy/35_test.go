package easy

import "testing"

func TestSearchInsert(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		out    int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 2, 1},
		{[]int{0, 1, 4, 5, 7, 12, 45}, 39, 6},
		{[]int{1, 2, 5, 6}, 3, 2},
		{[]int{10, 20, 30, 40, 50}, 35, 3},
		{[]int{}, 1, 0},
		{[]int{2}, 2, 0},
		{[]int{1}, 2, 1},
		{[]int{2}, 1, 0},
		{[]int{1, 4, 9}, 9, 2},
		{[]int{100, 222, 1111}, 221, 1},
	}
	for _, c := range cases {
		if res := searchInsert(c.nums, c.target); res != c.out {
			t.Errorf("searchInsert(%v, %d) = %d, expected %d", c.nums, c.target, res, c.out)
		}
	}

}
