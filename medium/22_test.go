package medium

import (
	"reflect"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	cases := []struct {
		in  int
		out []string
	}{
		{0, []string{}},
		{1, []string{"()"}},
		{2, []string{"()()", "(())"}},
		{3, []string{"()()()", "(()())", "((()))", "(())()", "()(())"}},
		{4, []string{
			"()()()()", // 1 1 1 1
			"(((())))", // 4
			"(()()())", // 4
			"((())())", // 4
			"(()(()))", // 4
			"((()()))", // 4
			"(())(())", // 2 2
			"((()))()", // 3 1
			"(()())()", // 3 1
			"()((()))", // 1 3
			"()(()())", // 1 3
			"()()(())", // 1 1 2
			"()(())()", // 1 2 1
			"(())()()", // 2 1 1
		}},
		{5, []string{
			"()()()()()", // 1 1 1 1 1

			"()(())(())", // 1 2 2
			"(())(())()", // 2 2 1
			"(())()(())", // 2 1 2

			"((()))(())", // 3 2
			"(()())(())", // 3 2

			"(())((()))", // 2 3
			"(())(()())", // 2 3

			"(()())()()", // 3 1 1
			"((()))()()", // 3 1 1
			"()(()())()", // 1 3 1
			"()((()))()", // 1 3 1
			"()()(()())", // 1 1 3
			"()()((()))", // 1 1 3

			"(())()()()", // 2 1 1 1
			"()(())()()", // 1 2 1 1
			"()()(())()", // 1 1 2 1
			"()()()(())", // 1 1 1 2

			"(((())))()", // 4 1
			"(()()())()", // 4 1
			"((()()))()", // 4 1
			"((())())()", // 4 1
			"(()(()))()", // 4 1

			"()(((())))", // 1 4
			"()(()()())", // 1 4
			"()((()()))", // 1 4
			"()((())())", // 1 4
			"()(()(()))", // 1 4

			"(()()()())", // 5
			"((((()))))", // 5
			"((()()()))", // 5
			"(((())()))", // 5
			"((()(())))", // 5
			"(((()())))", // 5
			"((())(()))", // 5
			"(((()))())", // 5
			"(()((())))", // 5
			"(()()(()))", // 5
			"(()(())())", // 5
			"((())()())", // 5
			"((()())())", // 5
			"(()(()()))", // 5
		}},
	}
	makeSet := func(ss []string) map[string]bool {
		set := make(map[string]bool)
		for _, s := range ss {
			set[s] = true
		}
		return set
	}
	for _, c := range cases {
		res := generateParentheses(c.in)
		if !reflect.DeepEqual(makeSet(res), makeSet(c.out)) {
			t.Errorf("generateParentheses(%d)=%v, expected %v", c.in, res, c.out)
		}

	}
}
