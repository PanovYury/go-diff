package diff

import (
	"slices"

	"github.com/PanovYury/go-diff/internal/lcs"
	"golang.org/x/exp/constraints"
)

type DiffResult[T comparable] struct {
	result bool
	item *T
}

func (dr *DiffResult[T]) Compare(dr2 DiffResult[T]) bool {
	return *dr.item == *dr2.item
}

func Diff[T constraints.Ordered](x, y *[]T) []DiffResult[T] {
	lcs := lcs.From(x, y)
	result := make([]DiffResult[T], 0, len(*x))
	for _, v := range *x {
		result = append(result, DiffResult[T]{
			result: slices.Index(lcs, &v) > -1,
			item: &v,
		})
	}
	return result
}
