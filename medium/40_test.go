package medium

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	var (
		makeBag = func(numsets [][]int) map[int]int {
			var res = make(map[int]int)
			for _, set := range numsets {
				for _, num := range set {
					if _, ok := res[num]; !ok {
						res[num] = 1
					} else {
						res[num]++
					}
				}
			}
			return res
		}
		testEqualBag = func(one, two [][]int) bool {
			return reflect.DeepEqual(makeBag(one), makeBag(two))
		}
		testEqualSize = func(one, two [][]int) bool {
			return len(one) == len(two)
		}
		testUniqueSetinSets = func(sets [][]int) bool { // test whether the set of number in the set of sets is unique
			for i := range sets {
				sort.Ints(sets[i])
			}
			for i := range sets {
				for j := range sets {
					if i == j {
						continue
					}
					if reflect.DeepEqual(sets[i], sets[j]) {
						return false
					}
				}
			}
			return true
		}
		testSumToTarget = func(sets [][]int, target int) bool {
			for _, set := range sets {
				var total = 0
				for _, num := range set {
					total += num
				}
				if total != target {
					return false
				}
			}
			return true
		}
	)
	cases := []struct {
		cand   []int
		target int
		res    [][]int
	}{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{
			[]int{1, 7},
			[]int{1, 2, 5},
			[]int{2, 6},
			[]int{1, 1, 6},
		}},
		{[]int{1}, 1, [][]int{
			[]int{1},
		}},
		{[]int{1, 1}, 1, [][]int{
			[]int{1},
		}},
	}

	for _, c := range cases {
		res := combinationSum2(c.cand, c.target)
		err := ""
		if !testEqualBag(res, c.res) || !testEqualSize(res, c.res) {
			err += fmt.Sprintf("size and member do not match to correct result.\n")

		}
		if !testUniqueSetinSets(res) {
			err += fmt.Sprintf("result set member is not unique: %v.\n", res)

		}
		if !testSumToTarget(res, c.target) {
			err += fmt.Sprintf("at least one set member (%v) does not sum up to target(%d).\n", res, c.target)

		}
		if len(err) != 0 {
			err += fmt.Sprintf("combinationSum2(%v,%d) = %v, expected %v", c.cand, c.target, res, c.res)
			t.Error(err)
		}
	}
}
