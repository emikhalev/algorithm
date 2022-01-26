// https://leetcode.com/problems/merge-intervals/
package merge_interval

import "sort"

func merge(intervals [][]int) [][]int {
	sortByStartPos(intervals)

	res := make([][]int, 0, len(intervals))
	interval := make([]int, 2)
	for i := 0; i < len(intervals); i++ {

		interval[0] = intervals[i][0]
		interval[1] = intervals[i][1]

		for ; i < len(intervals)-1 && interval[1] >= intervals[i+1][0]; i++ {
			interval[1] = max(interval[1], intervals[i+1][1])
		}

		res = append(res, []int{interval[0], interval[1]})
	}

	return res
}

func sortByStartPos(intervals [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
