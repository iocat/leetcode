package easy

import "math"

/* Problem 7:
 Given a 32-bit signed integer, reverse digits of an integer.

	Example 1:
		Input: 123
		Output:  321

	Example 2:
		Input: -123
		Output: -321

	Example 3:
		Input: 120
		Output: 21
Note:
Assume we are dealing with an environment which could only hold integers within
the 32-bit signed integer range. For the purpose of this problem, assume that
your function returns 0 when the reversed integer overflows
*/

func reverse(x int) int {
	var abs = func(num int32) (absolute float64, negative bool) { // returns the absolute value and whether the number is a negative
		if num < 0 {
			absolute = (float64)(-num)
			negative = true
			return
		}
		absolute = (float64)(num)
		negative = false
		return
	}
	var (
		val, neg = abs((int32)(x))
		reverse  float64           // the absolute value of the result
		step     = math.Log10(val) // the decimal value of the digit
	)

	for val > 0 {
		dec := (int32)(val) % 10 // get the last digit of the remaining
		val = val / 10           // remove the last digit from the remaining
		reverse += (float64)(dec) * math.Pow10((int)(step))
		step--
	}

	if reverse > math.MaxInt32 {
		return 0
	}

	if neg {
		return (int)(-(int32)(reverse))
	}
	return (int)((int32)(reverse))
}
