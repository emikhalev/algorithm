// https://leetcode.com/problems/permutations/
package permutations

func permute(nums []int) [][]int {
	return pm(nums, 0)
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
