// https://leetcode.com/problems/maximum-subarray/
// https://en.wikipedia.org/wiki/Maximum_subarray_problem
package oN_solution

func maxSubArray(nums []int) int {
	sum := 0
	maxSum := -2 << 31

	for i := 0; i < len(nums); i++ {
		sum = max(nums[i], sum+nums[i])
		maxSum = max(sum, maxSum)
	}

	return maxSum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
