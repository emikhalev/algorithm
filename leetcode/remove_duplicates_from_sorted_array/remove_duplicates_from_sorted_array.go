// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	found := make(map[int]struct{})
	l := 0
	j := 0

	for i := 0; i < len(nums); i++ {
		if isFound(found, nums, i) {
			j := max(i+1, j)
			for ; j < len(nums) && isFound(found, nums, j); j++ {
			}

			if j < len(nums) {
				swap(nums, i, j)
				i--
			}

			continue
		}
		found[nums[i]] = struct{}{}
		l++
	}

	return l
}

func isFound(found map[int]struct{}, nums []int, i int) bool {
	_, ok := found[nums[i]]
	return ok
}

func swap(nums []int, i, j int) {
	nextV := nums[j]
	nums[j] = nums[i]
	nums[i] = nextV
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
