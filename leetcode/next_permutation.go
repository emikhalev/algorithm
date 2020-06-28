// https://leetcode.com/problems/next-permutation/
package main

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i := len(nums) - 1
	for ; i > 0 && (nums[i] <= nums[i-1]); i-- {
	}

	if i > 0 {
		j := len(nums) - 1
		for ; j >= 0 && nums[i-1] >= nums[j]; j-- {
		}
		swap(nums, i-1, j)
	}

	reverse(nums, i)
}

func reverse(nums []int, start int) {
	i := start
	j := len(nums) - 1
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums []int, i, j int) {
	v := nums[i]
	nums[i] = nums[j]
	nums[j] = v
}
