package medium

var phone = [][]byte{
	[]byte{' '},                // #0
	[]byte{},                   // #1
	[]byte{'a', 'b', 'c'},      // #2
	[]byte{'d', 'e', 'f'},      // #3
	[]byte{'g', 'h', 'i'},      // #4
	[]byte{'j', 'k', 'l'},      // #5
	[]byte{'m', 'n', 'o'},      // #6
	[]byte{'p', 'q', 'r', 's'}, // #7
	[]byte{'t', 'u', 'v'},      // #8
	[]byte{'w', 'x', 'y', 'z'}, // #9
}

// Recursive Breadth First solution
// Worst case running time will be O(4^N) where N is the number of digits
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var letterComb func([]byte) []string
	letterComb = func(bytes []byte) []string {
		if len(bytes) == 1 {
			res := make([]string, 0, len(phone[bytes[0]-'0']))
			for _, char := range phone[bytes[0]-'0'] {
				res = append(res, string(char))
			}
			return res
		}
		var (
			charsToAppend = phone[bytes[len(bytes)-1]-'0'] // the list of character to append to all the previously generated strings
			oldSet        = letterComb(bytes[:len(bytes)-1])
			newSet        = make([]string, 0, len(oldSet)*len(charsToAppend))
		)
		for _, char := range charsToAppend {
			for _, str := range oldSet {
				newSet = append(newSet, str+string(char))
			}

		}
		return newSet
	}
	return letterComb([]byte(digits))
}
