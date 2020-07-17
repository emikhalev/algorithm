// https://en.wikipedia.org/wiki/Binary_heap

package heap

type BinaryHeapFn struct {
	data []int
	compareFn func(i, j int) int
}

func (h *BinaryHeapFn) Insert(v int) {
	h.data = append(h.data, v)
	h.siftUp()
}

func (h *BinaryHeapFn) Extract() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}

	rv := h.data[0]
	li := len(h.data)-1
	h.data[0] = h.data[li]
	h.data = h.data[:li]

	h.siftDown(0)
	return rv, true
}

func (h *BinaryHeapFn) siftUp() {
	c := len(h.data) - 1
	for ;c>0; {
		p := (c-1) / 2
		if h.compareFn(h.data[c], h.data[p]) < 0 {
			break
		}
		t := h.data[c]
		h.data[c] = h.data[p]
		h.data[p] = t
		c = p
	}
}

func (h *BinaryHeapFn) siftDown(i int) {
	for ;; {
		left := 2*i + 1
		right := 2*i + 2
		largest := i
		if left < len(h.data) && h.compareFn(h.data[left], h.data[largest]) > 0 {
			largest = left
		}
		if right < len(h.data) && h.compareFn(h.data[right], h.data[largest]) > 0 {
			largest = right
		}
		if largest == i {
			break
		}

		t := h.data[i]
		h.data[i] = h.data[largest]
		h.data[largest] = t

		i = largest
	}
}
