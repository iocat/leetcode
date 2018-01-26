package easy

/*
Given a sorted array and a target value, return the index if the target is
found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:
	Input: [1,3,5,6], 5
	Output: 2

Example 2:
	Input: [1,3,5,6], 2
	Output: 1

Example 3:
	Input: [1,3,5,6], 7
	Output: 4

Example 1:
	Input: [1,3,5,6], 0
	Output: 0
*/

// Running time is O(logN).
// Unexpectedly performing slower than a naive approach (O(N))
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	var searchFunc func(l, r int) int
	searchFunc = func(l, r int) int {
		if l >= r {
			if target <= nums[l] { // we use l in this case because in the recursive step r is moved over l (hence r is wrong)
				return l
			}
			return l + 1
		}
		var mid = int((l + r) / 2)
		if midval := nums[mid]; midval == target {
			return mid
		} else if midval > target {
			return searchFunc(l, mid-1) // could make r = l - 1 where l is expected to be the correct point, r is wrong
		} else {
			return searchFunc(mid+1, r) // could make l = r + 1 where l is expected to be the correct point, r is wrong
		}

	}

	return searchFunc(0, len(nums)-1)
}
