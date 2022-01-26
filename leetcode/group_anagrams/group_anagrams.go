// https://leetcode.com/problems/group-anagrams/
package group_anagrams

import (
	"crypto/sha1"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	hashes := getHashes(strs)

	indexes := make(map[[20]byte][]int)
	for i := range strs {
		h := hashes[i]
		if v, ok := indexes[h]; !ok {
			indexes[h] = []int{i}
		} else {
			indexes[h] = append(v, i)
		}
	}

	groups := make([][]string, len(indexes))
	i := 0
	for _, v := range indexes {
		groups[i] = make([]string, len(v))
		for j := 0; j < len(v); j++ {
			idx := v[j]
			groups[i][j] = strs[idx]
		}
		i++
	}

	return groups
}

// Count numbers of letters for each word O(N)
func getHashes(strs []string) [][20]byte {
	sorted := make([][]byte, len(strs))
	for i, v := range strs {
		sorted[i] = countSort(v)
	}

	hashes := make([][20]byte, len(sorted))
	for i, v := range sorted {
		hashes[i] = sha1.Sum(v)
	}

	return hashes
}

// O(N)
func countSort(strs string) []byte {
	// TODO: Using golang standard sort for now
	bStrs := []byte(strs)
	sort.Slice(bStrs, func(i, j int) bool { return bStrs[i] < bStrs[j] })
	return bStrs
}
