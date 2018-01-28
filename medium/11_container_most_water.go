package medium

import "math"

// Running time is O(N)
func maxArea(height []int) int {

	var (
		l, r    = 0, len(height) - 1
		maxFunc = func(x, y int) int {
			if x > y {
				return x
			}
			return y
		}
		minFunc = func(x, y int) int {
			if x < y {
				return x
			}
			return y
		}
	)
	var max = math.MinInt32
	for l != r {
		max = maxFunc(max, minFunc(height[l], height[r])*(r-l))
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return max
}

// Brute Force Solution
// Worst case running time O(N^2)
func maxArea0(height []int) int {
	var minFunc = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	var maxFunc = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	if len(height) == 2 {
		return minFunc(height[0], height[1])
	}

	var max = math.MinInt32

	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			max = maxFunc(max, minFunc(height[i], height[j])*(j-i))
		}
	}

	return max
}
