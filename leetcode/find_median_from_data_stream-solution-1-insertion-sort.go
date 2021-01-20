// https://leetcode.com/problems/find-median-from-data-stream/
package main

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
	if len(this.data) == 0 {
		this.data = append(this.data, num)
		return
	}

	// Binary Search
	l, r := 0, len(this.data)-1
	for l < r {
		m := l + (r-l)/2
		if this.data[m] < num {
			l = m + 1
			continue
		}
		r = m
	}

	if this.data[l] > num {
		l--
	}

	if l < 0 {
		this.data = append([]int{num}, this.data...)
		return
	}

	if l >= len(this.data)-1 {
		this.data = append(this.data, num)
		return
	}

	tmpData := make([]int, l+1, len(this.data)+1)
	copy(tmpData, this.data[0:l+1])
	tmpData = append(tmpData, num)
	if l+1 < len(this.data) {
		tmpData = append(tmpData, this.data[l+1:]...)
	}
	this.data = tmpData
}

func (this *MedianFinder) FindMedian() float64 {
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

/*
["MedianFinder","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian"]
[[],[-1],[],[-2],[],[-3],[],[-4],[],[-5],[]]
["MedianFinder","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian"]
[[],[6],[],[10],[],[2],[],[6],[],[5],[],[0],[],[6],[],[3],[],[1],[],[0],[],[0],[]]
*/
