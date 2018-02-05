package medium

import "testing"

func TestCanJump(t *testing.T) {
	cases := []struct {
		in  []int
		out bool
	}{
		{[]int{0, 0, 0, 1}, false},
		{[]int{1, 1}, true},
		{[]int{1, 0}, true},
		{[]int{1, 0, 3}, false},
		{[]int{1, 3, 3, 0, 0, 0, 0, 3}, false},
		{[]int{1, 3, 2, 0, 2, 0, 2, 3}, true},
		{[]int{29, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3}, true},
		{[]int{1, 20, 0, 10, 0, 0, 0, 00, 0, 0, 0, 0, 0, 0, 0, 0, 1}, true},
		{[]int{0, 11}, false},
	}
	for _, c := range cases {
		if res := canJump(c.in); res != c.out {
			t.Errorf("canJump(%v) = %v, expected %v", c.in, res, c.out)
		}
		if res := canJumpGreedy(c.in); res != c.out {
			t.Errorf("canJumpGreedy(%v) = %v, expected %v", c.in, res, c.out)
		}
	}
}
