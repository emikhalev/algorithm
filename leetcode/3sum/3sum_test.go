package _sum

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	cases := []struct {
		params []int
		exp    [][]int
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}

	for idx, tc := range cases {
		out := threeSum(tc.params)
		if !reflect.DeepEqual(out, tc.exp) {
			t.Errorf("case %d error: got %v, expected %v", idx, out, tc.exp)
		}
	}
}
