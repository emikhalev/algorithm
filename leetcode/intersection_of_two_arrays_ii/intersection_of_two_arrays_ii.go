// https://leetcode.com/problems/intersection-of-two-arrays-ii/
package intersection_of_two_arrays_ii

func intersect(nums1 []int, nums2 []int) []int {
	hm2 := slice2Map(nums2)
	res := make([]int, 0, min(len(nums1), len(nums2)))
	for i := 0; i < len(nums1); i++ {
		if cnt, ok := hm2[nums1[i]]; ok && cnt > 0 {
			res = append(res, nums1[i])
			hm2[nums1[i]] = cnt - 1
		}
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func slice2Map(nums []int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		v, ok := m[nums[i]]
		if !ok {
			v := 0
			m[nums[i]] = v
		}
		m[nums[i]] = v + 1
	}
	return m
}
