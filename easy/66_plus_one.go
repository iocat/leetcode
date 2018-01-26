package easy

/*
Given a non-negative integer represented as a non-empty array of digits, plus one to the integer.

You may assume the integer do not contain any leading zero, except the number 0 itself.

The digits are stored such that the most significant digit is at the head of the list.
*/
func plusOne(digits []int) []int {
	var po func(i int, overflow bool) []int

	po = func(i int, overflow bool) []int {
		if i == -1 {
			// prepend the result with one more digit
			return append([]int{1}, digits...)
		}
		if overflow {
			digits[i]++
			if digits[i] == 10 {
				digits[i] = 0
				return po(i-1, true)
			}
		}

		return digits
	}
	return po(len(digits)-1, true)
}
