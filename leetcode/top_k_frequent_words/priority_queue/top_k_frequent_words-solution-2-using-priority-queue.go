// https://leetcode.com/problems/top-k-frequent-words

package priority_queue

import (
	"strings"
)

func topKFrequent(words []string, k int) []string {
	counts := freqs(words)

	uniqWords := make([]string, 0, len(counts))
	for word, _ := range counts {
		uniqWords = append(uniqWords, word)
	}

	h := new(BinaryHeap)
	h.compareFn = func(w1, w2 string) int {
		cI := counts[w1]
		cJ := counts[w2]
		if cI != cJ {
			return cJ - cI
		}
		return strings.Compare(w1, w2)
	}

	for _, w := range uniqWords {
		h.Insert(w)
		if h.Size() > k {
			h.Extract()
		}
	}

	res := make([]string, 0, len(counts))
	for h.Size() > 0 {
		w, _ := h.Extract()
		res = append(res, w)
	}
	reverse(res)
	return res
}

func reverse(words []string) {
	for i := 0; i < len(words)/2; i++ {
		j := len(words) - 1 - i
		t := words[i]
		words[i] = words[j]
		words[j] = t
	}
}

func freqs(words []string) map[string]int {
	counts := make(map[string]int)
	for i := 0; i < len(words); i++ {
		word := words[i]
		cnt, _ := counts[word]
		cnt++
		counts[word] = cnt
	}
	return counts
}

type BinaryHeap struct {
	data      []string
	compareFn func(w1, w2 string) int
}

func (bh *BinaryHeap) Size() int {
	return len(bh.data)
}

func (bh *BinaryHeap) Insert(v string) {
	bh.data = append(bh.data, v)
	bh.siftUp()
}

func (bh *BinaryHeap) Extract() (string, bool) {
	if len(bh.data) == 0 {
		return "", false
	}

	rv := bh.data[0]
	li := len(bh.data) - 1
	bh.data[0] = bh.data[li]
	bh.data = bh.data[:li]

	bh.siftDown(0)

	return rv, true
}

func (bh *BinaryHeap) siftUp() {
	c := len(bh.data) - 1
	for c > 0 {
		p := (c - 1) / 2
		if bh.compareFn(bh.data[c], bh.data[p]) < 0 {
			break
		}
		t := bh.data[c]
		bh.data[c] = bh.data[p]
		bh.data[p] = t
		c = p
	}
}

func (bh *BinaryHeap) siftDown(i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		largest := i
		if left < len(bh.data) && bh.compareFn(bh.data[left], bh.data[largest]) > 0 {
			largest = left
		}
		if right < len(bh.data) && bh.compareFn(bh.data[right], bh.data[largest]) > 0 {
			largest = right
		}

		if largest == i {
			break
		}

		t := bh.data[i]
		bh.data[i] = bh.data[largest]
		bh.data[largest] = t

		i = largest
	}
}
