package medium

import (
	"sort"
	"strings"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {

	var sortElems = func(strSet [][]string) {
		for i := range strSet {
			sort.Strings(strSet[i])
		}
	}
	// make a sort function for the string set given that the
	// string set are individually sorted
	var makeSortFn = func(set [][]string) func(int, int) bool {
		return func(a, b int) bool {
			if len(set[a]) == 0 || len(set[b]) == 0 {
				panic("expect a non-empty string set")
			}
			return strings.Compare(set[a][0], set[b][0]) < 0
		}
	}

	// check if the 2 sorted sets are equal
	var equalSortedSet = func(one, two [][]string) bool {
		if len(one) != len(two) {
			return false
		}
		for i, strs := range one {
			if len(strs) != len(two[i]) {
				return false
			}
			for j, str := range strs {
				if two[i][j] != str {
					return false
				}
			}
		}
		return true
	}

	cases := []struct {
		in  []string
		out [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				[]string{"ate", "eat", "tea"},
				[]string{"nat", "tan"},
				[]string{"bat"},
			}},
		{[]string{"ca", "ac", "mad", "sak", "kas", "dma", "tmm"},
			[][]string{
				[]string{"ca", "ac"},
				[]string{"mad", "dma"},
				[]string{"sak", "kas"},
				[]string{"tmm"},
			}},
		{[]string{"lmll", "asds", "wm", "dsas", "mw", "tks", "skt", "lllm", "llml"},
			[][]string{
				[]string{"lmll", "lllm", "llml"},
				[]string{"asds", "dsas"},
				[]string{"wm", "mw"},
				[]string{"tks", "skt"},
			}},
		{[]string{"lol"}, [][]string{[]string{"lol"}}},
		{[]string{}, [][]string{}},
		{[]string{"mm", "mm"}, [][]string{[]string{"mm", "mm"}}},
	}
	for _, c := range cases {
		res := groupAnagrams(c.in)
		// Sort the output
		sortElems(c.out)
		sort.Slice(c.out, makeSortFn(c.out))
		// Sort the expected output
		sortElems(res)
		sort.Slice(res, makeSortFn(res))

		// compare
		if !equalSortedSet(res, c.out) {
			t.Errorf("groupAnagrams(%v) = %v, expected %v", c.in, res, c.out)
		}

	}
}
