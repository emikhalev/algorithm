package find_target_indices_after_sorting_array

import (
	"reflect"
	"testing"
)

func Test_targetIndicesSolutionV2(t *testing.T) {
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
		out := targetIndicesSolutionV2(tc.nums, tc.target)
		if !reflect.DeepEqual(out, tc.exp) {
			t.Errorf("case %d error: \n-- got: %v\n-- expected: %v", idx, out, tc.exp)
		}
	}
}

func Benchmark_targetIndicesSolutionV2(b *testing.B) {
	const (
		target = 7
		cnt    = 30
	)
	benchCase := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		benchCase[i] = target
	}

	for i := 0; i < b.N; i++ {
		targetIndicesSolutionV2(benchCase, target)
	}
}
