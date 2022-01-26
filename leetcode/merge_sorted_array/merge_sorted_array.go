// https://leetcode.com/problems/merge-sorted-array/
package merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	iW := n + m - 1
	for i >= 0 && j >= 0 {
		a := nums1[i]
		b := nums2[j]

		if a > b {
			nums1[iW] = a
			i--
		}

		if a <= b {
			nums1[iW] = b
			j--
		}

		iW--
	}

	for ; i >= 0; i-- {
		nums1[iW] = nums1[i]
		iW--
	}

	for ; j >= 0; j-- {
		nums1[iW] = nums2[j]
		iW--
	}
}
