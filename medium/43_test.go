package medium

import "testing"

func TestMultiply(t *testing.T) {
	cases := []struct {
		in  [2]string
		out string
	}{
		{[2]string{"0", "0"}, "0"},
		{[2]string{"12", "20"}, "240"},
		{[2]string{"1", "0"}, "0"},
		{[2]string{"0", "1"}, "0"},
		{[2]string{"9", "9"}, "81"},
		{[2]string{"12321328131293021391203921039210392109401294021490219421", "0"}, "0"},
		{[2]string{"20", "1"}, "20"},
	}
	for _, c := range cases {
		if res := multiply(c.in[0], c.in[1]); res != c.out {
			t.Errorf("multiply(%s,%s)=%s, expected %s", c.in[0], c.in[1], res, c.out)
		}
	}
}
