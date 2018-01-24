package easy

import "testing"

func TestStringInString(t *testing.T) {
	cases := []struct {
		haystack string
		needle   string
		index    int
	}{
		{"", "", 0},
		{"a", "", 0},
		{"", "a", -1},
		{"a", "aaaa", -1},
		{"aa", "a", 0},
		{"aa", "ab", -1},
		{"abcd", "c", 2},
		{"12abc", "1", 0},
		{"112233akq", "3ak", 5},
		{"111111111111", "111", 0},
		{"11223312", "12", 1},
	}
	for _, c := range cases {
		if res := strStr(c.haystack, c.needle); res != c.index {
			t.Errorf("strStr(\"%s\",\"%s\") = %d, expected %d", c.haystack, c.needle, res, c.index)
		}
	}
}
