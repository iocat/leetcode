package easy

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return reverse(x) == x // reverse implemented in problem #7
}
