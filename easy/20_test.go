package easy

import "testing"

func TestValidParentheses(t *testing.T) {
	cases := []struct {
		in  string
		out bool
	}{
		{"((((", false},
		{"", true},
		{"(((())))", true},
		{"([})", false},
		{"([{}()([]){}])", true},
		{"([{{{}}}])", true},
		{"[[{{{{]]", false},
		{"()()", true},
		{"()[(()){}]", true},
	}
	for _, c := range cases {
		if isValid(c.in) != c.out {
			t.Errorf("isValid(\"%s\") = %v, expected %v", c.in, !c.out, c.out)
		}
	}
}
