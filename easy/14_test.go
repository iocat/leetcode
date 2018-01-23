package easy

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		in  []string
		out string
	}{
		{
			[]string{},
			"",
		},
		{
			nil,
			"",
		},
		{
			[]string{""},
			"",
		},
		{
			[]string{"a", "aa"},
			"a",
		},
		{
			[]string{"abc", "abcd", "a"},
			"a",
		},
		{
			[]string{"", "aslkdnsal"},
			"",
		},
		{
			[]string{"aaaa", "bbbb"},
			"",
		},
		{
			[]string{"bbbb", "bbcd", "bbbbwd"},
			"bb",
		},
	}

	for _, c := range cases {
		if res := longestCommonPrefix(c.in); res != c.out {
			t.Errorf("longestCommonPrefix(%v)=%s, expected %s", c.in, res, c.out)
		}
	}
}
