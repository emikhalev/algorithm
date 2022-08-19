package fibonacci_number

import (
	"testing"
)

func Test_fib(t *testing.T) {
	cases := []struct {
		num      int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{30, 832040},
	}

	for idx, tc := range cases {
		res := fib(tc.num)
		if tc.expected != res {
			t.Errorf("case %d. expecting %d, got %d", idx, tc.expected, res)
		}
	}
}
