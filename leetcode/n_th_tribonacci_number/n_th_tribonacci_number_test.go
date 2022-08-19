package n_th_tribonacci_number

import (
	"testing"
)

func Test_tribonacci(t *testing.T) {
	cases := []struct {
		num      int
		expected int
	}{
		{4, 4},
		{25, 1389537},
	}

	for idx, tc := range cases {
		res := tribonacci(tc.num)
		if tc.expected != res {
			t.Errorf("case %d. expecting %d, got %d", idx, tc.expected, res)
		}
	}
}
