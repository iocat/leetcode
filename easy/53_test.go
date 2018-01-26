package easy

import "testing"

func TestMaxSubarray(t *testing.T) {
	cases := []struct {
		nums  []int
		total int
	}{
		{[]int{}, 0},
		{[]int{-1}, -1},
		{[]int{-1, -2}, -1},
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{10, -100, 20, 10, -10}, 30},
		{[]int{1, 2, -3, 3, 3, 3}, 9},
	}
	for _, c := range cases {
		if res := maxSubArray(c.nums); res != c.total {
			t.Errorf("maxSubArray(%v) = %d, expected %d", c.nums, res, c.total)
		}
	}
}
