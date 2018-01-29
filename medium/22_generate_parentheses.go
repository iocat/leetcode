package medium

func generateParentheses(n int) []string {
	if n == 0 {
		return []string{}
	}

	results := make([]string, 0)

	var backtrack func(results *[]string, progress string, open, close int, max int)
	backtrack = func(results *[]string, progress string, open, close int, max int) {
		if len(progress) == max*2 {
			*results = append(*results, progress)
			return
		}
		if open < max {
			backtrack(results, progress+"(", open+1, close, max)
		}
		if close < open {
			backtrack(results, progress+")", open, close+1, max)
		}
	}
	backtrack(&results, "", 0, 0, n)

	return results
}
