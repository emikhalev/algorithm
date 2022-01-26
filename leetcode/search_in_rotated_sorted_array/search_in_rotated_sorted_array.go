// https://leetcode.com/problems/search-in-rotated-sorted-array/
package search_in_rotated_sorted_array

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	minIdx := binaryFindMinIdx(nums, 0, len(nums)-1)
	return binarySearch(nums, target, 0, len(nums)-1, minIdx)
}

func binarySearch(nums []int, target int, l, r, shift int) int {
	if r < l {
		return -1
	}

	mid := l + (r-l)/2
	midVal := getVal(nums, mid, shift)

	if midVal == target {
		return (mid + shift) % len(nums)
	}

	if midVal > target {
		return binarySearch(nums, target, l, mid-1, shift)
	}
	return binarySearch(nums, target, mid+1, r, shift)
}

func getVal(nums []int, i, shift int) int {
	idx := (i + shift) % len(nums)
	return nums[idx]
}

func binaryFindMinIdx(nums []int, l, r int) int {
	if r < l {
		return 0
	}

	mid := l + (r-l)/2
	if mid+1 >= len(nums) {
		return 0
	}

	if nums[mid] > nums[mid+1] {
		return mid + 1
	}
	if nums[mid] < nums[l] {
		return binaryFindMinIdx(nums, l, mid-1)
	} else {
		return binaryFindMinIdx(nums, mid+1, r)
	}

	return 0
}
