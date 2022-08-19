package min_cost_climbing_stairs

import "testing"

func Test_minCostClimbingStairs(t *testing.T) {
	cases := []struct {
		num      []int
		expected int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}

	for idx, tc := range cases {
		res := minCostClimbingStairs(tc.num)
		if tc.expected != res {
			t.Errorf("case %d. expecting %d, got %d", idx, tc.expected, res)
		}
	}
}
