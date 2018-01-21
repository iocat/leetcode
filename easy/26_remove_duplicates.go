package easy

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
