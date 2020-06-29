// https://leetcode.com/problems/two-sum/
package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		b := target - a

		if idx, ok := m[b]; ok && idx != i {
			return []int{i, idx}
		}

	}
	return []int{}
}
