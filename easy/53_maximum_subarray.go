package easy

import "math"

/*
Problem 53:
Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
the contiguous subarray [4,-1,2,1] has the largest sum = 6.
*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		total      int
		totalSoFar = math.MinInt32
	)
	for _, val := range nums {
		total += val
		if total > totalSoFar {
			totalSoFar = total
		}
		if total <= 0 {
			total = 0
		}
	}
	return totalSoFar
}

// DP solution
func maxSubArray2(nums []int) int {
	var max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	if len(nums) == 0 {
		return 0
	}
	var maxSofar = math.MinInt32
	var maxEndHere int
	for _, v := range nums {
		maxEndHere = max(v, maxEndHere+v)
		maxSofar = max(maxSofar, maxEndHere)
	}
	return maxSofar
}
