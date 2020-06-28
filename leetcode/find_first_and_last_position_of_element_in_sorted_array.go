// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
package main

func searchRange(nums []int, target int) []int {
	r := []int{-1, -1}

	if len(nums) == 0 {
		return r
	}

	/*if len(nums) == 1 {
	    return []int{0, 0}
	}*/

	r[0] = getStart(nums, target)
	if r[0] == -1 {
		return r
	}
	r[1] = getEnd(nums, target, r[0])

	return r
}

func getStart(nums []int, target int) int {
	if nums[0] == target {
		return 0
	}
	n := len(nums)
	idx := 0
	for s := n / 2; s >= 1; s /= 2 {
		for idx+s < n && (nums[idx+s] < target || (nums[idx+s] == target && nums[idx+s-1] < target)) {
			idx += s
		}

		if nums[idx] == target && nums[idx-1] < target {
			return idx
		}
	}

	return -1
}

func getEnd(nums []int, target int, startPos int) int {

	if nums[len(nums)-1] == target {
		return len(nums) - 1
	}

	n := len(nums)
	idx := startPos
	for s := (n - idx) / 2; s >= 1; s /= 2 {
		for idx+s < n && (nums[idx+s] == target || (nums[idx+s] > target && nums[idx+s-1] == target)) {
			idx += s
		}

		if nums[idx] > target && nums[idx-1] == target {
			return idx - 1
		}
	}
	return -1
}
