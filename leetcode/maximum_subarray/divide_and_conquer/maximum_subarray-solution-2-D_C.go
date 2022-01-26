// https://leetcode.com/problems/maximum-subarray/
// https://en.wikipedia.org/wiki/Maximum_subarray_problem
// Divide and Conquer approach
package divide_and_conquer

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	mid := (n - 1) / 2

	ans := max(maxSubArray(nums[:mid+1]), maxSubArray(nums[mid+1:]))

	sum := 0
	pLeft := nums[mid]
	for i := mid; i >= 0; i-- {
		sum += nums[i]
		pLeft = max(pLeft, sum)
	}

	pRight := nums[mid+1]
	sum = 0
	for i := mid + 1; i < n; i++ {
		sum += nums[i]
		pRight = max(pRight, sum)
	}

	return max(ans, pLeft+pRight)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
