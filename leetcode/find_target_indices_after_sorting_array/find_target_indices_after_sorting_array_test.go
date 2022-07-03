package find_target_indices_after_sorting_array

import (
	"reflect"
	"testing"
)

func Test_targetIndices(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		exp    []int
	}{
		{
			nums:   []int{1, 2, 5, 2, 3},
			target: 2,
			exp:    []int{1, 2},
		},
		{
			nums:   []int{1, 2, 5, 2, 3},
			target: 3,
			exp:    []int{3},
		},
		{
			nums:   []int{1, 2, 5, 2, 3},
			target: 5,
			exp:    []int{4},
		},
	}

	for idx, tc := range cases {
		out := targetIndices(tc.nums, tc.target)
		if !reflect.DeepEqual(out, tc.exp) {
			t.Errorf("case %d error: \n-- got: %v\n-- expected: %v", idx, out, tc.exp)
		}
	}
}
