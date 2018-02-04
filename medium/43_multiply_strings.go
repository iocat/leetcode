package medium

import (
	"bytes"
	"fmt"
)

func byteToDigit(b byte) int {
	return int(b - '0')
}

var numDict = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func digitToByte(d int) byte {
	return numDict[d]
}

func reverse(num []byte) []byte {
	var ret = make([]byte, len(num))
	for i := len(num) - 1; i >= 0; i-- {
		ret[len(num)-1-i] = num[i]
	}
	return ret
}

func multiply(num1, num2 string) string {
	if l1, l2 := len(num1), len(num2); l1 == l2 && l1 == 0 {
		return ""
	}
	var (
		fst, snd = reverse([]byte(num1)), reverse([]byte(num2))
		ret      = make([]int, len(num1)+len(num2))
	)

	for i, d1 := range fst {
		for j, d2 := range snd {
			writeTo := i + j
			overflown := byteToDigit(d1) * byteToDigit(d2)
			for overflown > 0 {
				ret[writeTo] += overflown

				if written := ret[writeTo]; written >= 10 {
					overflown = int(written / 10)
					ret[writeTo] = written % 10
					writeTo++
				} else {
					overflown = 0 // stop the cascading additions
				}
			}

		}
	}

	// Construct the result
	for i := len(ret) - 1; i >= 0; i-- { // remove trailing zero
		if ret[i] == 0 {
			if i == 0 {
				return "0"
			}
			ret = ret[0:i]
		} else {
			break
		}
	}
	var res bytes.Buffer
	for i := len(ret) - 1; i >= 0; i-- {
		fmt.Println(ret, i)
		res.WriteByte(digitToByte(ret[i]))
	}
	return res.String()
}
