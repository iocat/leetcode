package medium

/*

Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

For example:
A = [2,3,1,1,4], return true.

A = [3,2,1,0,4], return false.

*/
// The running time is O(N^2)
func canJumpTo(nums []int, ind int, cache []*bool) bool {
	if ind > len(nums)-1 || ind < 0 {
		return false
	} else if cache[ind] != nil {
		return *(cache[ind])
	} else if ind == 0 {
		return true
	}
	for i := ind - 1; i >= 0; i-- {
		tmp := canJumpTo(nums, i, cache) && (nums[i] >= ind-i)
		if tmp {
			cache[ind] = new(bool)
			*(cache[ind]) = true // cache
			return true
		}
	}
	// can't reach to ind from here
	cache[ind] = new(bool)
	*(cache[ind]) = false
	return false
}

func canJumpGreedy(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	var leftMostGoodIndex = len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= leftMostGoodIndex {
			leftMostGoodIndex = i
		}
	}
	return leftMostGoodIndex == 0
}

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	cache := make([]*bool, len(nums))
	cache[0] = new(bool)
	*(cache[0]) = true
	return canJumpTo(nums, len(nums)-1, cache)
}
