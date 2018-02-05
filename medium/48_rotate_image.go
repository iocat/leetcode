package medium

func rotate4ways(m [][]int, i, j int) {
	swap := func(a, b *int) {
		*a, *b = *b, *a
	}
	n := len(m)
	swap(&m[j][n-1-i], &m[i][j])
	swap(&m[n-1-i][n-1-j], &m[i][j])
	swap(&m[n-1-j][i], &m[i][j])
}

// depict a virtual wing
func rotate(m [][]int) {
	var (
		row  = 0
		l, r = 0, len(m) - 2
	)
	for l <= r {
		for i := l; i <= r; i++ {
			rotate4ways(m, row, i)
		}
		row++
		l++
		r--
	}
}
