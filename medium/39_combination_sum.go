package medium

import (
	"reflect"
	"sort"
)

/*

Given a set of candidate numbers (C) (without duplicates) and a target number
(T), find all unique combinations in C where the candidate numbers sums to T.

The same repeated number may be chosen from C unlimited number of times.

Note:
	All numbers (including target) will be positive integers.
	The solution set must not contain duplicate combinations.

For example, given candidate set [2, 3, 6, 7] and target 7,
A solution set is:
	[
	[7],
	[2, 2, 3]
	]

*/

// First, we sort the candidate numbers
// Second, use backtracking to generate all the sets
// Third, eliminate all the duplicated set
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	// Sort the candidates
	sort.Ints(candidates) // O(NlogN)

	// Exlude all candidates > target
	for i, num := range candidates { // O(N)
		if num > target {
			candidates = candidates[:i]
		}
	}

	var res = make([][]int, 0)

	var backtracking func(
		res *[][]int,
		curSet []int, // the accumulated set
		curSum int, // the accumulated sum of curSet
	)
	backtracking = func(res *[][]int, curSet []int, curSum int) {
		if curSum == target { // previous step adds up to the target
			dupSet := make([]int, len(curSet))
			copy(dupSet, curSet)
			*res = append(*res, dupSet)
		}
		for _, cand := range candidates {
			if cand > target-curSum { // no longer be able to add up
				break
			}
			backtracking(res, append(curSet, cand), curSum+cand)
		}
	}
	backtracking(&res, []int{}, 0)

	// Remove duplicated set
	for _, set := range res {
		sort.Ints(set)
	}
	var toEliminate = make([]bool, len(res))
	for i := 0; i < len(res)-1; i++ {
		if toEliminate[i] {
			continue
		}
		for j := i + 1; j < len(res); j++ {
			if toEliminate[j] {
				continue
			}
			if reflect.DeepEqual(res[i], res[j]) {
				toEliminate[j] = true
			}
		}
	}

	count := 0
	for i, eliminate := range toEliminate {
		if eliminate {
			res = append(res[:i-count], res[i-count+1:]...)
			count++
		}
	}

	return res
}
