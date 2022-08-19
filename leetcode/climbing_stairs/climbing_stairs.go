// https://leetcode.com/problems/climbing-stairs/
package climbing_stairs

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
