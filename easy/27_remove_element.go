package easy

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
