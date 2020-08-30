// https://leetcode.com/problems/top-k-frequent-words

package main

import (
	"sort"
	"strings"
)

func topKFrequent(words []string, k int) []string {
	counts := freqs(words)

	res := make([]string, 0, len(counts))
	for word, _ := range counts {
		res = append(res, word)
	}

	sort.Slice(res, func(i, j int) bool {
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
