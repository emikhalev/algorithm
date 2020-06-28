// https://leetcode.com/problems/product-of-array-except-self/
package main

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	k := 1
	for i := 0; i < len(nums); i++ {
		res[i] = k
		k = k * nums[i]
	}

	k = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * k
		k = k * nums[i]
	}

	return res
}
