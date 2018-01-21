package easy

import (
	"testing"
)

func TestReverseInteger(t *testing.T) {
	var cases = []struct {
		in int
		ou int
	}{
		{12, 21},
		{1, 1},
		{120, 21},
		{-120, -21},
		{11, 11},
		{1234, 4321},
		{10, 1},
		{1534236469, 0},
		{-2147483648, 0},
	}

	for _, c := range cases {
		if res := reverse(c.in); res != c.ou {
			t.Errorf("reverse(%d): expected %d, received %d", c.in, c.ou, res)
		}
	}
}
