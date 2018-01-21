package easy

import "testing"

func TestIsPalindromeNumber(t *testing.T) {
	var cases = []struct {
		x          int
		palindrome bool
	}{
		{0, true},
		{1, true},
		{4, true},
		{9, true},
		{10, false},
		{0, true},
		{-100, false},
		{-29, false},
		{-101, false},
		{121, true},
		{1111, true},
		{1221, true},
		{100021, false},
		{1011, false},
	}
	for _, c := range cases {
		if isPalindrome(c.x) != c.palindrome {
			t.Errorf("isPalindrome(%d): expected %v, received %v", c.x, c.palindrome, !c.palindrome)
		}
	}
}
