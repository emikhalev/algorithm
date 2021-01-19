// https://leetcode.com/submissions/detail/445082440/

package main

import "sort"

func firstMissingPositive(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	expectedNum := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			continue
		}

		if nums[i] != expectedNum {
			break
		}

		for ; i < len(nums) && nums[i] == expectedNum; i++ {
		}
		i--
		expectedNum++
	}

	return expectedNum
}

// 1, 2, 0 -> 0, 1, 2
// n=0, exp=0, c; n=1, exp=0, exp=2(n+1), c; n=2, exp=2, exp=3(++); c; eoc; ret exp(3)

// 3, 4, -1, 1 -> -1, 1, 3, 4 -> 1
// n=-1, c; n=1, exp=2; n=3, exp=2, break; return exp(2)

// Trivial:
// 1. sort - N*logN CPU, O(N) - space complexity for merge sort, qsort
// 2. iterate - find first positive, after that find first missing positive

// Test cases
// 1. simple: [0, 1, 2, 3, 5], [34, 36, 50], [-50, -30, -3, -2, -1, 0, 1, 2, 3, 5]
// 2. missing two or more positives: [0, 1, 2, 3, 5, 8, 12]
// 3. no-zero: [[1, 2, 3, 5]]
// 4. no-zero and negatives: [-1, -2, -4, 1, 2, 3, 5]
// 5. empty: []
// 6. no-positives: [-5, -4, -3, -2, -1]
// 7. no-positives and zero: [-5, -4, -3, -2, -1, 0]
// 8. no-positives, zero and missing negatives: [-10, -8, -5, -4, -3, -2, -1, 0]
// 9. no missing positives: 2^31-1 [1<<31-3, 1<<31-2, 1<<31-1] -> [2147483645, 2147483646, 2147483647]
// 10. [0]
// 11. [1]
// 12. [2]
// 13. [-2]
// 14. Repeatable: [0,2,2,1,1]
