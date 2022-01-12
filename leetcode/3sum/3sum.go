// https://leetcode.com/problems/3sum/
package _sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	r := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		if i > 0 && nums[i-1] == a {
			continue
		}

		bc := twoSums(nums[i+1:], -a)
		for j := 0; j < len(bc); j++ {
			r = append(r, []int{a, bc[j][0], bc[j][1]})
		}
	}

	return r
}

func twoSums(nums []int, target int) [][]int {
	res := make([][]int, 0)
	l, r := 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]

		if sum > target || (r < len(nums)-1 && nums[r] == nums[r+1]) {
			r--
			continue
		}
		if sum < target || (l > 0 && nums[l-1] == nums[l]) {
			l++
			continue
		}

		if sum == target {
			res = append(res, []int{nums[l], nums[r]})
			l++
			r--
		}
	}
	return res
}
