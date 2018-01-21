package main

import "fmt"

func main() {
	var in = []int{0, 0, 1, 1, 1, 1, 2, 4, 4, 4, 5, 5, 9, 19}
	newL := removeDuplicates(in)
	fmt.Println("Remove Duplicate", in[:newL])
}

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
