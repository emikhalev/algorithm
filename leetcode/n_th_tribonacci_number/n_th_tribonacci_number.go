// https://leetcode.com/problems/n-th-tribonacci-number/
package n_th_tribonacci_number

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n < 3 {
		return 1
	}

	nums := make([]int, n+1)
	nums[0] = 0
	nums[1] = 1
	nums[2] = 1

	for i := 0; i < n-2; i++ {
		nums[i+3] = nums[i] + nums[i+1] + nums[i+2]
	}

	return nums[n]
}
