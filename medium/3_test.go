package medium

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		s string
		l int
	}{
		{"abcabcabc", 3},
		{"abcdef", 6},
		{"tylkklyytttt", 4},
		{"anqldddddkkaabcdefg", 7},
		{"kkk", 1},
		{"k", 1},
		{"", 0},
		{"llvvmm", 2},
	}
	for _, c := range cases {
		if res := lengthOfLongestSubstring(c.s); res != c.l {
			t.Errorf("lengthOfLongestSubstring(%s)=%d, expected %d", c.s, res, c.l)
		}
	}
}
