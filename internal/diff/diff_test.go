package diff

import "testing"

func TestDiff(t *testing.T) {
	x := []int32{1, 2, 3}
	y := []int32{1, 2, 3}
	diff := []DiffResult[int32] {
		{
			result: true,
			item: &x[0],
		},
		{
			result: true,
			item: &x[1],
		},
		{
			result: true,
			item: &x[2],
		},
	}

	_diff := Diff(&x, &y)

	if len(diff) != len(_diff) {
		t.Fatalf("Incorrect diff size: %v. Expected: %v", len(_diff), len(diff))
	}

	for i := range _diff {
		if !_diff[i].Compare(diff[i]) {
			t.Fatalf("Incorrect value in index %v: %v != %v", i, *_diff[i].item, *diff[i].item)
		}
	}
}