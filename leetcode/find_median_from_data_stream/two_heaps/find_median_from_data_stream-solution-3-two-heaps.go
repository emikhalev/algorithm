// https://leetcode.com/problems/find-median-from-data-stream/

package two_heaps

type MedianFinder struct {
	qLow, qHigh *PriorityQueue
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		qLow:  NewQueue(func(i, j int) bool { return i > j }),
		qHigh: NewQueue(func(i, j int) bool { return i < j }),
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.qLow.Push(num)

	v, _ := this.qLow.Top()
	this.qHigh.Push(v)
	this.qLow.Pop()

	if this.qLow.Len() < this.qHigh.Len() {
		v, _ := this.qHigh.Top()
		this.qLow.Push(v)
		this.qHigh.Pop()
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.qLow.Len() > this.qHigh.Len() {
		v, _ := this.qLow.Top()
		return float64(v)
	}
	vl, _ := this.qLow.Top()
	vh, _ := this.qHigh.Top()
	return float64(vl+vh) / 2
}

type PriorityQueue struct {
	compareFn func(i, j int) bool
	data      []int
}

func NewQueue(compareFn func(i, j int) bool) *PriorityQueue {
	return &PriorityQueue{
		compareFn: compareFn,
		data:      make([]int, 0),
	}
}

func (q *PriorityQueue) Len() int {
	return len(q.data)
}

func (q *PriorityQueue) Top() (int, bool) {
	if len(q.data) > 0 {
		return q.data[0], true
	}
	return 0, false
}

func (q *PriorityQueue) Push(num int) {
	q.data = append(q.data, num)
	q.siftUp()
}

func (q *PriorityQueue) Pop() (int, bool) {
	if len(q.data) == 0 {
		return 0, false
	}

	v := q.data[0]
	q.data[0] = q.data[len(q.data)-1]
	q.data = q.data[:len(q.data)-1]
	q.siftDown()

	return v, true
}

func (q *PriorityQueue) siftUp() {
	cIdx := len(q.data) - 1 // current index
	pIdx := (cIdx - 1) / 2

	for cIdx > 0 && q.compareFn(q.data[cIdx], q.data[pIdx]) {
		// xor swap (works for integers)
		q.data[cIdx] = q.data[cIdx] ^ q.data[pIdx]
		q.data[pIdx] = q.data[pIdx] ^ q.data[cIdx]
		q.data[cIdx] = q.data[cIdx] ^ q.data[pIdx]

		// update indexes
		cIdx = pIdx
		pIdx = (cIdx - 1) / 2
	}
}

func (q *PriorityQueue) siftDown() {
	if len(q.data) == 0 {
		return
	}

	cIdx := 0
	for {
		// left and right indexes
		lIdx := cIdx*2 + 1
		rIdx := lIdx + 1

		// find max
		maxIdx := cIdx
		if lIdx < len(q.data) && q.compareFn(q.data[lIdx], q.data[maxIdx]) {
			maxIdx = lIdx
		}
		if rIdx < len(q.data) && q.compareFn(q.data[rIdx], q.data[maxIdx]) {
			maxIdx = rIdx
		}

		// swap current with max
		if cIdx == maxIdx {
			break
		}

		// xor swap (works for integers)
		q.data[cIdx] = q.data[cIdx] ^ q.data[maxIdx]
		q.data[maxIdx] = q.data[maxIdx] ^ q.data[cIdx]
		q.data[cIdx] = q.data[cIdx] ^ q.data[maxIdx]

		// update current index
		cIdx = maxIdx
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
