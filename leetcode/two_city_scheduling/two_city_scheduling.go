// https://leetcode.com/problems/two-city-scheduling/
package two_city_scheduling

import (
	"sort"
)

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs,
		func(i, j int) bool { return costs[i][1]-costs[i][0] > costs[j][1]-costs[j][0] })

	total := 0
	for i := 0; i < len(costs); i++ {
		if i < len(costs)/2 {
			total += costs[i][0]
		} else {
			total += costs[i][1]
		}
	}

	return total
}
