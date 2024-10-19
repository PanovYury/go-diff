package lcs

import (
	"slices"

	"github.com/PanovYury/go-diff/internal/matrix"
	"golang.org/x/exp/constraints"
)

func From[T constraints.Ordered](x, y *[]T) []*T {
	matrix := matrix.From(x, y)
	lcs := make([]*T, 0, len(*x))

	for i, j := len(*x), len(*y); i > 0 && j > 0; {
		if (*x)[i-1] == (*y)[j-1] {
			lcs = append(lcs, &(*x)[i-1])
			i -= 1
			j -= 1
		} else if matrix[i-1][j] == matrix[i][j] {
			i -= 1
		} else {
			j -= 1
		}
	}

	slices.Reverse(lcs)
	return lcs
}