// https://leetcode.com/problems/find-target-indices-after-sorting-array/
package find_target_indices_after_sorting_array

import (
	"sort"
)

func targetIndices(nums []int, target int) []int {
	sort.Sort(sort.IntSlice(nums))

	// Find elements
	l, r := 0, len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}

	// Collect all indices
	res := make([]int, 0)
	for ; l < len(nums) && nums[l] == target; l++ {
		res = append(res, l)
	}
	return res
}
