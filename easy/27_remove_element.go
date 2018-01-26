package easy

/*
Problem 27:
Given an array and a value, remove all instances of that value in-place and
return the new length.

Do not allocate extra space for another array, you must do this by modifying
the input array in-place with O(1) extra memory.

The order of elements can be changed. It doesn't matter what you leave beyond
the new length.

Example:

Given nums = [3,2,2,3], val = 3,

Your function should return length = 2, with the first two elements of nums
being 2.

*/

func removeElements(nums []int, val int) int {
	var (
		head = 0
		end  = 0
	)
	if len(nums) == 0 {
		return 0
	}
loop:
	for head < len(nums) {
		for nums[head] == val {
			head++
			if head == len(nums) { // found the value and reached the end at the same time
				break loop
			}
		}
		nums[end] = nums[head]
		head, end = head+1, end+1 // advance both the head and the end
	}
	return end
}
