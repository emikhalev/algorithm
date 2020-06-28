// https://leetcode.com/problems/container-with-most-water
package main

func maxArea(height []int) int {
	n := len(height)
	l, r := 0, n-1
	maxV := 0

	for l < r {
		h := min(height[l], height[r])
		v := h * (r - l)
		if maxV < v {
			maxV = v
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxV
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
