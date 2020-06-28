// https://leetcode.com/problems/permutations-ii/
package main

import (
	"fmt"
)

func permuteUnique(nums []int) [][]int {
	ps := pm(nums, 0)
	return rmDups(ps)
}

func rmDups(inS [][]int) [][]int {
	outS := make([][]int, 0, len(inS))

	m := make(map[string]struct{})
	for i := 0; i < len(inS); i++ {
		k := fmt.Sprintf("%x", inS[i])
		if _, ok := m[k]; ok {
			continue
		}
		outS = append(outS, inS[i])
		m[k] = struct{}{}
	}

	return outS
}

func pm(nums []int, pos int) [][]int {
	r := make([][]int, 0)
	tr := make([]int, len(nums))
	copy(tr, nums)
	r = append(r, tr)
	i := pos
	for j := pos + 1; j < len(nums); j++ {
		tr := make([]int, len(nums))
		copy(tr, nums)
		swap(tr, i, j)
		r = append(r, tr)
	}

	res := make([][]int, 0, 2)

	if pos < len(nums)-2 {
		for i := 0; i < len(r); i++ {
			rc := make([]int, len(r[i]))
			copy(rc, r[i])
			pr := pm(rc, pos+1)
			res = append(res, pr...)
		}
	} else {
		res = append(res, r...)
	}

	return res
}

func swap(sl []int, i, j int) {
	t := sl[i]
	sl[i] = sl[j]
	sl[j] = t
}
