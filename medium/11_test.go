package medium

import "testing"

func getTestMaxAreaCases() []struct {
	height []int
	area   int
} {
	return []struct {
		height []int
		area   int
	}{
		{[]int{1, 2, 3, 4, 5}, 6},
		{[]int{1, 2}, 1},
		{[]int{2, 3, 2, 2}, 6},
		{[]int{1, 1}, 1},
		{[]int{4, 2, 5}, 8},
		{[]int{2, 3, 4, 3, 2}, 8},
	}
}

func TestMaxArea(t *testing.T) {
	cases := getTestMaxAreaCases()
	for _, c := range cases {
		if res := maxArea(c.height); res != c.area {
			t.Errorf("maxArea(%v) = %d, expected %d", c.height, res, c.area)
		}
	}
}

func TestMaxArea0(t *testing.T) {
	cases := getTestMaxAreaCases()
	for _, c := range cases {
		if res := maxArea0(c.height); res != c.area {
			t.Errorf("maxArea0(%v) = %d, expected %d", c.height, res, c.area)
		}
	}
}
