package medium

import (
	"sort"
)

func combinationSum2BackTrack(selectFrom, cur []int, target, start int, ret *[][]int) {
	if target < 0 {
		return
	} else if target == 0 {
		var tmp = make([]int, len(cur))
		copy(tmp, cur)
		*ret = append(*ret, tmp)
	} else {
		for i := start; i < len(selectFrom) && target >= selectFrom[i]; i++ {
			l := len(cur)
			if i > start && selectFrom[i] == selectFrom[i-1] {
				continue // avoid generation of duplicated sets (say 1 1 2) can results in 1(first 1) 2 and 1 (2nd 1) 2
			}

			cur = append(cur, selectFrom[i])
			combinationSum2BackTrack(selectFrom, cur, target-selectFrom[i], i+1, ret)
			cur = cur[:l]
		}
	}
}

// The idea is to check every combinations
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res = make([][]int, 0)
	combinationSum2BackTrack(candidates, make([]int, 0), target, 0, &res)
	return res
}
