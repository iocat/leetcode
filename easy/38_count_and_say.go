package easy

import (
	"bytes"
)

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
			count = 0          // count the occurence of the head
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
