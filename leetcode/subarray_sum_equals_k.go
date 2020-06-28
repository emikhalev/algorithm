// https://leetcode.com/problems/subarray-sum-equals-k/
package main

func subarraySum(nums []int, k int) int {
	sum := make(map[int]int)
	s := 0
	sum[0] = 1
	cnt := 0
	for i := 0; i < len(nums); i++ {
		s = s + nums[i]

		if v, ok := sum[s-k]; ok {
			cnt += v
		}

		if _, ok := sum[s]; ok {
			sum[s] = sum[s] + 1
		} else {
			sum[s] = 1
		}
	}

	return cnt
}
