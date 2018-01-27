package medium

import "math"

/*
Problem 3:
Given a string, find the length of the longest substring without repeating characters.

Examples:
	Given "abcabcbb", the answer is "abc", which the length is 3.
	Given "bbbbb", the answer is "b", with the length of 1.
	Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

// Brute Force: taking too long to complete
func lengthOfLongestSubstring0(s string) int {
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 1
	}
	var lengthSoFar = math.MinInt32
	for l := 0; l < len(s)-1; l++ {
	RLOOP:
		for r := l; r < len(s); r++ {
			sub := s[l : r+1]
			chars := make(map[byte]bool)
			for _, c := range []byte(sub) {
				if !chars[c] {
					chars[c] = true
				} else {
					continue RLOOP
				}
			}
			if len(sub) > lengthSoFar {
				lengthSoFar = len(sub)
			}
		}
	}
	return lengthSoFar
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var (
		sb      = []byte(s)
		cm      = make(map[byte]int)
		head    = 0
		max     = math.MinInt32
		maxFunc = func(x, y int) int {
			if x > y {
				return x
			}
			return y
		}
	)
	for tail, v := range sb {
		if val, ok := cm[v]; ok {
			head = maxFunc(head, val+1) // only allows advancing forward. Otherwise the head may revert back to the first few characters
		}
		cm[v] = tail
		max = maxFunc(max, tail-head+1)
	}
	return max
}
