// https://leetcode.com/problems/find-peak-element/
package main

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	l := 0
	r := len(nums) - 1

	for r >= l {
		mid := l + (r-l)/2

		lv := getValue(nums, mid-1)
		vv := getValue(nums, mid)
		rv := getValue(nums, mid+1)

		if lv <= vv && vv >= rv {
			return mid
		}

		if vv < lv {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

func getValue(nums []int, idx int) int {
	if idx < 0 || idx >= len(nums) {
		return -1 << 32
	}

	return nums[idx]
}
