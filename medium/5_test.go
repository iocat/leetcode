package medium

import "testing"

func getPalindromeTestCases() []struct {
	in, out string
} {
	return []struct {
		in  string
		out string
	}{
		{"abccba", "abccba"},
		{"abbc", "bb"},
		{"_", "_"},
		{"abaccccaba", "abaccccaba"},
		{"", ""},
		{"aaab1baaa", "aaab1baaa"},
		{"@!#$$#@!", "#$$#"},
		{"a11aa11a", "a11aa11a"},
		{"abaclwkeqqqqqq", "qqqqqq"},
		{"aaaa", "aaaa"},
	}
}

func TestLongestPalindromeSubstr(t *testing.T) {
	cases := getPalindromeTestCases()
	for _, c := range cases {
		if res := longestPalindrome(c.in); res != c.out {
			t.Errorf("longestPalindrome(%s)=%s, expected %s", c.in, res, c.out)
		}
	}
}

func TestLengthOfLongestSubstring0(t *testing.T) {
	cases := getPalindromeTestCases()
	for _, c := range cases {
		if res := longestPalindrome0(c.in); res != c.out {
			t.Errorf("longestPalindrome0(%s)=%s, expected %s", c.in, res, c.out)
		}
	}
}
