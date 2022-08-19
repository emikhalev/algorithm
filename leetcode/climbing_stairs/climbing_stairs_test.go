package climbing_stairs

import "testing"

func Test_climbStairs(t *testing.T) {
	cases := []struct {
		num      int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{45, 1836311903},
		{10, 89},
	}

	for idx, tc := range cases {
		res := climbStairs(tc.num)
		if tc.expected != res {
			t.Errorf("case %d. expecting %d, got %d", idx, tc.expected, res)
		}
	}
}
