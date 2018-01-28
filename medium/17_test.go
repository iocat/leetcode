package medium

import (
	"reflect"
	"testing"
)

func makeSet(ss []string) map[string]bool {
	var set = make(map[string]bool)
	for _, s := range ss {
		set[s] = true
	}
	return set
}
func TestLetterCombination(t *testing.T) {
	cases := []struct {
		digits string
		res    []string
	}{
		{"", []string{}},
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"2", []string{"a", "b", "c"}},
		{"67", []string{"mp", "mq", "mr", "ms", "np", "nq", "nr", "ns", "op", "oq", "or", "os"}},
		{"98", []string{"wt", "wu", "wv", "xt", "xu", "xv", "yt", "yu", "yv", "zt", "zu", "zv"}},
		{"55", []string{"jj", "jk", "jl", "kj", "kk", "kl", "lj", "lk", "ll"}},
	}
	for _, c := range cases {
		if res := letterCombinations(c.digits); !reflect.DeepEqual(makeSet(res), makeSet(c.res)) {
			t.Errorf("letterCombinations(%s) = %v, expected %v", c.digits, res, c.res)
		}
	}
}
