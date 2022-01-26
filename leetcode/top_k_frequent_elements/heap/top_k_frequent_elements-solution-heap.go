// https://leetcode.com/problems/top-k-frequent-elements

// Based on Min Binary Heap: https://en.wikipedia.org/wiki/Binary_heap
package heap

func topKFrequent(nums []int, k int) []int {
	freqs := make(map[int]int)

	// O(N)
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		cnt, _ := freqs[n]
		cnt++
		freqs[n] = cnt
	}

	// O(N log k)
	h := Heap{
		compareFn: func(i, j int) int {
			return freqs[j] - freqs[i]
		},
	}

	for num, _ := range freqs {
		h.Insert(num)
		if h.Size() > k {
			h.Extract()
		}
	}

	// O(k log k)
	ans := make([]int, h.Size())
	idx := 0
	for h.Size() > 0 {
		ans[idx], _ = h.Extract()
		idx++
	}
	return ans
}

type Heap struct {
	data      []int
	compareFn func(i, j int) int
}

func (h *Heap) Size() int {
	return len(h.data)
}

func (h *Heap) Insert(v int) {
	h.data = append(h.data, v)
	h.siftUp(len(h.data) - 1)
}

func (h *Heap) Extract() (v int, ok bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	v = h.data[0]
	idx := len(h.data) - 1
	h.data[0] = h.data[idx]
	h.data = h.data[:idx]
	h.siftDown(0)
	return v, true
}

func (h *Heap) siftUp(i int) {
	c := i
	for c > 0 {
		p := (c - 1) / 2
		if h.compareFn(h.data[c], h.data[p]) < 0 {
			break
		}
		t := h.data[c]
		h.data[c] = h.data[p]
		h.data[p] = t
		c = p
	}
}

func (h *Heap) siftDown(i int) {
	for {
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
