package medium

/*
Given a collection of distinct numbers, return all possible permutations.

For example,
[1,2,3] have the following permutations:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/

// The idea is to continuously swap the element to the front. The right side of
// the list is guaranteed to be unique. The elements which are swapped into
// position is fixed, then we start swapping the rest of the list. Each time we
// can no longer swap we add the temporary set to the result.
func permute(nums []int) [][]int {
	var backtrack func(res *[][]int, tmp []int, ind int)
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
		for i := ind; i < len(tmp); i++ {
			swap(&tmp[ind], &tmp[i])
			backtrack(res, tmp, ind+1)
			swap(&tmp[ind], &tmp[i]) // undo the swapping operation
		}
	}
	res := make([][]int, 0)
	backtrack(&res, nums, 0)
	return res
}
