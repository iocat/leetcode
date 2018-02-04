package medium

/*
Given a collection of numbers that might contain duplicates, return all possible unique permutations.

For example,
[1,1,2] have the following unique permutations:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/

// Solution: like a permutation generation problem given the given numbers are distinct (#46),
// each backtracking step needs a set to cache all the swapped value and make sure
// not to swap that same value again to the index we're swapping against
//
// Note To Self: For backtracking, enunumarate the result sets using a recursion tree
// Then figure out a way to eliminate all the duplicated results (in this case,
// a hashmap is required to store the swapped value)
func permuteUnique(nums []int) [][]int {
	var backtrack func(
		res *[][]int, // the accumulated result
		tmp []int, // the temporary result
		ind int, // the index we want to swap to the back
	)
	var swap = func(a, b *int) {
		*a, *b = *b, *a
	}
	backtrack = func(res *[][]int, tmp []int, ind int) {
		if ind == len(tmp)-1 {
			var ns = make([]int, len(tmp))
			copy(ns, tmp)
			*res = append(*res, ns)
			return
		}
		swappedSet := make(map[int]bool)
		for i := ind; i < len(tmp); i++ {
			if swappedSet[tmp[i]] {
				continue
			} else {
				swappedSet[tmp[i]] = true
			}
			swap(&tmp[ind], &tmp[i])
			backtrack(res, tmp, ind+1)
			swap(&tmp[ind], &tmp[i]) // undo the swapping operation
		}
	}
	res := make([][]int, 0)
	backtrack(&res, nums, 0)
	return res
}
