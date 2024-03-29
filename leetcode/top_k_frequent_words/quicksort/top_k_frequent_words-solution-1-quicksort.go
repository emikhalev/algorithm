// https://leetcode.com/problems/top-k-frequent-words

package quicksort

import (
	"strings"
)

func topKFrequent(words []string, k int) []string {
	counts := freqs(words)

	res := make([]string, 0, len(counts))
	for word, _ := range counts {
		res = append(res, word)
	}

	quicksort(res, 0, len(res)-1, func(i, j int) bool {
		cI := counts[res[i]]
		cJ := counts[res[j]]
		if cI != cJ {
			return cI > cJ
		}
		return strings.Compare(res[i], res[j]) < 0
	})

	return res[0:k]
}

func freqs(words []string) map[string]int {
	counts := make(map[string]int)
	for i := 0; i < len(words); i++ {
		word := words[i]
		cnt, _ := counts[word]
		cnt++
		counts[word] = cnt
	}
	return counts
}

func quicksort(nums []string, left, right int, compareFn func(i, j int) bool) {
	if left < right {
		pivotIndex := partition(nums, left, right, compareFn)
		quicksort(nums, left, pivotIndex-1, compareFn)
		quicksort(nums, pivotIndex+1, right, compareFn)
	}
}

func partition(nums []string, left, right int, compareFn func(i, j int) bool) (pivotIndex int) {
	i := left
	for j := left; j <= right; j++ {
		if compareFn(j, right) {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, right)
	return i
}

func swap(nums []string, i, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}
