package easy

import "strings"

/*
Given a string s consists of upper/lower-case alphabets and empty space
characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.
*/
func lengthOfLastWord(s string) int {
	res := strings.Split(strings.Trim(s, " "), " ")
	if len(res) == 0 {
		return 0
	}
	return len(res[len(res)-1])
}
