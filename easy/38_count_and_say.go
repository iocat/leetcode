package easy

import "bytes"

/*
The count-and-say sequence is the sequence of integers with the first five terms as following:

1.     1
2.     11
3.     21
4.     1211
5.     111221

1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.
Given an integer n, generate the nth term of the count-and-say sequence.

Note: Each term of the sequence of integers will be represented as a string.

Example 1:
	Input: 1
	Output: "1"

Example 2:
	Input: 4
	Output: "1211"
*/

func countAndSay(n int) string {
	var (
		i       = 1
		res     = []byte("1")
		byteMap = [...]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	)
	if n < 1 {
		panic("invalid parameter expected >= 1")
	}
	for i != n {
		var (
			newS  bytes.Buffer // the new string
			head  = 0          // the parsing progress
			count = 0          // the number of the head's occurences
		)

		for head < len(res) {
			for head+count < len(res) && res[head+count] == res[head] {
				count++
			}
			newS.WriteByte(byteMap[count])
			newS.WriteByte(res[head])
			head = head + count // advance the head
			count = 0           // reset the counter
		}
		res = newS.Bytes()
		i++
	}
	return string(res)
}
