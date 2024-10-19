package lcs

import (
	"slices"
	"testing"
)

type TestCase[T any] struct {
	x []T
	y []T
	expected []T
}

func parse[T any](arr []*T) []T {
	result := make([]T, len(arr))
	for i, v := range arr {
		result[i] = *v
	}
	return result
}

func TestLcs(t *testing.T) {
	cases := []TestCase[int32] {
		{
			x: []int32{1, 2, 3},
			y: []int32{1, 2, 3},
			expected: []int32{1, 2, 3},
		},
		{
			x: []int32{1, 2, 3},
			y: []int32{4, 5, 6},
			expected: []int32{},
		},
		{
			x: []int32{1, 2, 3},
			y: []int32{1, 5, 3},
			expected: []int32{1, 3},
		},
		{
			x: []int32{1, 2, 3},
			y: []int32{4, 2, 6},
			expected: []int32{2},
		},
		{
			x: []int32{1, 2, 3},
			y: []int32{4, 2, 3},
			expected: []int32{2, 3},
		},
	}

	for _, c := range cases {
		actual := parse(From(&c.x, &c.y))

		if !slices.Equal(c.expected, actual) {
			t.Errorf("Expected %v, but got %v", c.expected, actual)
		}
	}
}

func Benchmark100(b *testing.B) {
	data := make([]string, 100)
	for i := range data {
		data[i] = "12345678901234567890123456789012345678901234567890"
	}

	b.ResetTimer()

	From(&data, &data)

	b.ReportAllocs()
}

func Benchmark1000(b *testing.B) {
	data := make([]string, 1000)
	for i := range data {
		data[i] = "12345678901234567890123456789012345678901234567890"
	}

	b.ResetTimer()

	From(&data, &data)
}

func Benchmark10000(b *testing.B) {
	data := make([]string, 10_000)
	for i := range data {
		data[i] = "12345678901234567890123456789012345678901234567890"
	}

	b.ResetTimer()

	From(&data, &data)
}

func BenchmarkMillion(b *testing.B) {
	data := make([]string, 1_000_000)
	for i := range data {
		data[i] = "12345678901234567890123456789012345678901234567890"
	}

	b.ResetTimer()

	From(&data, &data)
}
