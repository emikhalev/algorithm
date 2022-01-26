// https://leetcode.com/problems/interval-list-intersections/
package interval_list_intersections

func intervalIntersection(A [][]int, B [][]int) [][]int {
	r := make([][]int, 0)

	i := 0
	j := 0
	for i < len(A) && j < len(B) {

		aStart, aEnd := A[i][0], A[i][1]
		bStart, bEnd := B[j][0], B[j][1]
		maxStart := max(aStart, bStart)
		minEnd := min(aEnd, bEnd)

		if maxStart <= minEnd {
			r = append(r, []int{maxStart, minEnd})
		}

		if aEnd > bEnd {
			j++
		} else {
			i++
		}
	}

	return r
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
