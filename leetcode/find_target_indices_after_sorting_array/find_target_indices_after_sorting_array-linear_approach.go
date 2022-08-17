package find_target_indices_after_sorting_array

import (
	"sort"
)

func targetIndicesLinear(nums []int, target int) []int {
	sort.Sort(sort.IntSlice(nums))

	res := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			res = append(res, i)
		}
	}

	return res
}
