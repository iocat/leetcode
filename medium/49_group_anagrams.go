package medium

import "math"

/*
Given an array of strings, group anagrams together.

For example, given: ["eat", "tea", "tan", "ate", "nat", "bat"],
Return:
	[
	["ate", "eat","tea"],
	["nat","tan"],
	["bat"]
	]
	Note: All inputs will be in lower-case.
*/

// The expected running time is O(N)
// The worst space complexity is O(|N|), worst case happens O(N) when the
// given input strings are not anagrams to one another

// The idea is to create a character set for each output set of anagrams.
// We use an array of size math.MaxUint8 to store the number of occurences of
// each character, then we make that array a key of a set which maps to the
// anagrams list. Whenever a string with a same set of character reoccurs,
//  that string is appended to the result set
func groupAnagrams(strs []string) [][]string {
	var set = make(map[[math.MaxUint8]int][]string)
	for _, str := range strs {
		var charSet = [math.MaxUint8]int{}
		for _, letter := range []byte(str) {
			charSet[letter]++
		}
		if _, ok := set[charSet]; !ok {
			set[charSet] = []string{str}
		} else {
			set[charSet] = append(set[charSet], str)
		}
	}
	res := make([][]string, 0, len(set))
	for _, v := range set {
		res = append(res, v)
	}
	return res
}
