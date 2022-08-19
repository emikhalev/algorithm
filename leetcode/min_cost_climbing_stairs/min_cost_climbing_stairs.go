// https://leetcode.com/problems/min-cost-climbing-stairs/
package min_cost_climbing_stairs

func minCostClimbingStairs(cost []int) int {
	n := len(cost)

	dp := make([]int, n+1)
	dp[n] = 0
	dp[n-1] = cost[n-1]

	for i := n - 2; i >= 0; i-- {
		dp[i] = cost[i] + min(dp[i+1], dp[i+2])
	}

	return min(dp[0], dp[1])
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
