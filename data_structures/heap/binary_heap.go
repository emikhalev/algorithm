// https://en.wikipedia.org/wiki/Binary_heap

package heap

/*
May be stored in:
1) An array (started with 0 index)
left = 2*i + 1
right = 2*i + 2
parent = trunc((i-1)/2)
where i - index of current node in array, root node index is zero

2) Data structure with pointers to child nodes

Operations: Insert, Extract
*/

type BinaryHeap struct {
	data []int
}

/*
Insert - add an element to a heap

1. Add the element to the bottom level of the heap at the most left.
2. Compare the added element with its parent; if they are in the correct order, stop.
3. If not, swap the element with its parent and return to the previous step
*/
func (h *BinaryHeap) Insert(v int) {
	h.data = append(h.data, v)
	h.siftUp()
}

/*
Extract - deleting the root from the heap

1. Replace the root of the heap with the last element on the last level.
2. Compare the new root with its children; if they are in the correct order, stop.
3. If not, swap the element with one of its children and return to the previous step. (Swap with its smaller child in a min-heap and its larger child in a max-heap.)
*/
func (h *BinaryHeap) Extract() (int, bool) {
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

// sift-up operation (after insert)
// also known as up-heap, bubble-up, percolate-up, sift-up, trickle-up, swim-up, heapify-up, or cascade-up
// Time complexity: O(log n)
/*
1. Compare the added element with its parent; if they are in the correct order, stop.
2. If not, swap the element with its parent and return to the previous step
*/
func (h *BinaryHeap) siftUp() {
	c := len(h.data) - 1
	for ;c>0; {
		p := (c-1) / 2
		if h.data[c] < h.data[p] {
			break
		}
		t := h.data[c]
		h.data[c] = h.data[p]
		h.data[p] = t
		c = p
	}
}

// sift-down operation (after extract)
// also known as down-heap, bubble-down, percolate-down, sift-down, sink-down, trickle down, heapify-down, cascade-down, and extract-min/max
// Time complexity: O(log n)
/*
1. Compare the new root with its children; if they are in the correct order, stop.
2. If not, swap the element with one of its children and return to the previous step.
	(Swap with its smaller child in a min-heap and its larger child in a max-heap.)
 */
func (h *BinaryHeap) siftDown(i int) {
	for ;; {
		left := 2*i + 1
		right := 2*i + 2
		largest := i
		if left < len(h.data) && h.data[left] > h.data[largest] {
			largest = left
		}
		if right < len(h.data) && h.data[right] > h.data[largest] {
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
