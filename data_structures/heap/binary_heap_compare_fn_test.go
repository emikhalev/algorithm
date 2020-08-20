package heap

import "testing"

var (
	compareFnMaxHeap = func(i, j int) int {
		return i - j
	}

	compareFnMinHeap = func(i, j int) int {
		return j - i
	}
)

func TestBinaryHeapFnMaxHeap(t *testing.T) {
	h := new(BinaryHeapFn)
	h.compareFn = compareFnMaxHeap

	h.Insert(4)
	h.Insert(5)
	h.Insert(1)
	h.Insert(10)
	h.Insert(3)
	h.Insert(1)
	h.Extract()
	h.Insert(5)
	h.Extract()
	h.Insert(7)
	h.Insert(2)

	expected := []int{7, 5, 4, 3, 2, 1, 1}
	for i, n := range expected {
		hn, _ := h.Extract()
		if n != hn {
			t.Errorf("expecting %d, got %d (test case number %d)\n", n, hn, i)
		}
	}
}

func TestBinaryHeapFnMinHeap(t *testing.T) {
	h := new(BinaryHeapFn)
	h.compareFn = compareFnMinHeap

	h.Insert(4)
	h.Insert(5)
	h.Insert(1)
	h.Insert(10)
	h.Insert(3)
	h.Insert(1)
	h.Extract()
	h.Insert(5)
	h.Extract()
	h.Insert(7)
	h.Insert(2)

	expected := []int{2, 3, 4, 5, 5, 7, 10}
	for i, n := range expected {
		hn, _ := h.Extract()
		if n != hn {
			t.Errorf("expecting %d, got %d (test case number %d)\n", n, hn, i)
		}
	}
}
