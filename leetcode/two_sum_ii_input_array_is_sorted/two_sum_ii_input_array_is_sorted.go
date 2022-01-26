// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
package two_sum_ii_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1

	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}
		if sum < target {
			i++
		}
		if sum > target {
			j--
		}
	}
	return []int{-1, -1}
}
