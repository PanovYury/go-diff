package matrix

import "golang.org/x/exp/constraints"

type Matrix = map[int]map[int]int32

func From[T constraints.Ordered](x, y *[]T) Matrix {
	matrix := make(map[int]map[int]int32)
	for i := range len(*x) + 1 {
		matrix[i] = make(map[int]int32)
	}

	for i := 1; i < len(*x) + 1; i++ {
		for j := 1; j < len(*y) + 1; j++ {
			if (*x)[i - 1] == (*y)[j - 1] {
				matrix[i][j] = matrix[i - 1][j - 1] + 1
			} else {
				matrix[i][j] = max(matrix[i - 1][j], matrix[i][j - 1])
			}
		}
	}
	
	return matrix
}
