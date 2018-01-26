package easy

import "bytes"

// Problem 14:
// Write a function to find the longest common prefix string amongst an array
// of strings.

// This algorithm starts by finding the length of the shortest string (minLen).
// Then, it examines all the character from index 0 -> minLen in each string looking for
// a single differing character at the same index. For each same character in the string set,
// it is added to the intermediate result. Once the differring character is found, the algorithm
// returns the constructed string so far.
// The runtime of this algorithm is O(n*m) where n is the number of input strings
// and m is the length of the minimal string.
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	var (
		minLen = len(strs[0]) // the length of minimum string
	)

	for i := 1; i < len(strs); i++ { //  look for the minimum length
		if tmp := len(strs[i]); tmp < minLen {
			minLen = tmp
		}
	}

	if minLen == 0 {
		return ""
	}

	var res bytes.Buffer // the result

	for i := 0; i < minLen; i++ {
		var ref = strs[0][i] // the character of the first string for occurence check
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != ref { // a differing character occurs at strs[j]
				return res.String() // result is the constructed string so far
			}
		}
		res.WriteByte(ref)
	}

	return res.String()
}
