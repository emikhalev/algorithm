// https://leetcode.com/problems/find-median-from-data-stream/
/*
From LeetCode solution:
-----------

Complexity Analysis

Time complexity: O(n\log n) + O(1) \simeq O(n\log n)O(nlogn)+O(1)≃O(nlogn).

Adding a number takes amortized O(1)O(1) time for a container with an efficient resizing scheme.
Finding the median is primarily dependent on the sorting that takes place. This takes O(n\log n)O(nlogn) time for a standard comparative sort.
Space complexity: O(n)O(n) linear space to hold input in a container. No extra space other than that needed (since sorting can usually be done in-place).
*/
package main

import "sort"

type MedianFinder struct {
	data []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		data: make([]int, 0, 16),
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.data = append(this.data, num)
}

func (this *MedianFinder) FindMedian() float64 {
	sort.Sort(sort.IntSlice(this.data))
	n := len(this.data)
	if n%2 == 1 {
		return float64(this.data[n/2])
	}

	return float64((this.data[n/2-1] + this.data[n/2])) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
