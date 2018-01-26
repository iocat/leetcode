package easy

/*
Problem 28:
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if
needle is not part of haystack.

Example 1:
	Input: haystack = "hello", needle = "ll"
	Output: 2


Example 2:
	Input: haystack = "aaaaa", needle = "bba"
	Output: -1

*/

func strStr(haystack string, needle string) int {
	switch { // precondition check
	case len(needle) == len(haystack):
		if needle == haystack {
			return 0
		}
	case len(haystack) == 0 && len(needle) != 0,
		len(needle) > len(haystack):
		return -1
	case len(needle) == 0:
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] { // potentially encounter a needle
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		} else {
			continue
		}
	}

	return -1
}
