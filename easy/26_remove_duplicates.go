package easy

/*
Problem 26:
Given a sorted array, remove the duplicates in-place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example:
	Given nums = [1,1,2],

	Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.
	It doesn't matter what you leave beyond the new length.

*/

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	var (
		head      = 0
		toReplace = 0
	)

	for head < len(nums) {
		counter := head + 1
		for counter < len(nums) && nums[counter] == nums[head] {
			counter++
		}
		nums[toReplace] = nums[head]
		toReplace++
		head = counter
	}
	return toReplace
}
