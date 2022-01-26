// https://leetcode.com/problems/intersection-of-two-arrays/
package intersection_of_two_arrays

func intersection(nums1 []int, nums2 []int) []int {
	hm2 := slice2Map(nums2)
	res := make([]int, 0, min(len(nums1), len(nums2)))
	for i := 0; i < len(nums1); i++ {
		if _, ok := hm2[nums1[i]]; ok {
			res = append(res, nums1[i])
			delete(hm2, nums1[i])
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

func slice2Map(nums []int) map[int]struct{} {
	m := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = struct{}{}
	}
	return m
}
