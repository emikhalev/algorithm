// https://leetcode.com/problems/subsets/
package main

func subsets(nums []int) [][]int {
	return getSubsets(nums, []int{}, 0)
}

func getSubsets(nums []int, subset []int, pos int) [][]int {
	res := make([][]int, 0)
	if pos == len(nums) {
		res = append(res, subset)
		return res
	}

	// Not include current char into subset
	s1 := make([]int, len(subset))
	copy(s1, subset)
	r1 := getSubsets(nums, s1, pos+1)

	// Include current char into subset
	s2 := make([]int, len(subset)+1)
	s2[len(s2)-1] = nums[pos]
	copy(s2, subset)
	r2 := getSubsets(nums, s2, pos+1)

	res = append(res, r1...)
	res = append(res, r2...)

	return res
}
