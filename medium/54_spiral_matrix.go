package medium

const (
	up = iota
	Down
	Left
	Right
)

// The idea is to gradually cut the outer layer of the matrix til the matrix is
// empty as we go element by element
func spiralOrderRecursive(matrix [][]int, dir int, ret *[]int) {
	if len(matrix) == 0 {
		return
	}

	switch dir {
	case up:
		for i := len(matrix) - 1; i >= 0; i-- {
			row := &matrix[i]
			*ret = append(*ret, (*row)[0])
			*row = (*row)[1:]
			if i == 0 && len(*row) == 0 {
				matrix = matrix[0:0]
			}
		}
		spiralOrderRecursive(matrix, Right, ret)
	case Down:
		for i := 0; i < len(matrix); i++ {
			row := &matrix[i]
			*ret = append(*ret, (*row)[len(*row)-1])
			*row = (*row)[:len(*row)-1]
			if i == len(matrix)-1 && len(*row) == 0 {
				matrix = matrix[0:0]
			}
		}
		spiralOrderRecursive(matrix, Left, ret)
	case Left:
		lastRow := &matrix[len(matrix)-1]
		*ret = append(*ret, (*lastRow)[len(*lastRow)-1]) // append the last row element
		*lastRow = (*lastRow)[:len(*lastRow)-1]
		if len(*lastRow) == 0 {
			matrix = matrix[:len(matrix)-1] // remove the last row
			spiralOrderRecursive(matrix, up, ret)
		} else {
			spiralOrderRecursive(matrix, dir, ret)
		}
	case Right:
		firstRow := &matrix[0]
		*ret = append(*ret, (*firstRow)[0])
		*firstRow = (*firstRow)[1:]
		if len(*firstRow) == 0 {
			matrix = matrix[1:] // cut the first row and move toward Down
			spiralOrderRecursive(matrix, Down, ret)
		} else {
			spiralOrderRecursive(matrix, dir, ret)
		}
	}
}

// Running time is O(n)
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	in := make([][]int, len(matrix))
	copy(in, matrix)
	var ret = make([]int, 0, len(in)*len(in[0]))
	spiralOrderRecursive(in, Right, &ret)
	return ret
}
