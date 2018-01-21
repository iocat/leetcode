package easy

import "testing"

func TestCountAndSay(t *testing.T) {
	var cases = []struct {
		n   int
		val string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
		{6, "312211"},
		{7, "13112221"},
		{8, "1113213211"},
		{9, "31131211131221"},
		{10, "13211311123113112211"},
		{11, "11131221133112132113212221"},
		{12, "3113112221232112111312211312113211"},
	}
	for _, c := range cases {
		if res := countAndSay(c.n); res != c.val {
			t.Errorf("error: countAndSay(%d) returns %s, expected %s", c.n, res, c.val)
		}
	}
}
