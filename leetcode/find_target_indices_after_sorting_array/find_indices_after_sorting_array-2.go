// https://leetcode.com/problems/find-target-indices-after-sorting-array/
package find_target_indices_after_sorting_array

import "sort"

func targetIndicesSolutionV2(nums []int, target int) []int {
	sort.Sort(sort.IntSlice(nums))

	// Find leftmost element
	l, r := 0, len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	leftIdx := l

	// Find rightmost element
	l, r = 0, len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	rightIdx := r - 1

	cnt := rightIdx - leftIdx + 1
	if cnt <=0 {
		return nil
	}

	// Collect all indices
	res := make([]int, cnt)
	for idx:=leftIdx;idx<=rightIdx;idx++ {
		res[idx-leftIdx] = idx
	}

	return res
}

