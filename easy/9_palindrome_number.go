package easy

// Problem 9:
// Determine whether an integer is a palindrome. Do this without extra space.

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return reverse(x) == x // reverse implemented in problem #7
}
