// https://leetcode.com/problems/insert-delete-getrandom-o1
package insert_delete_getrandom_o1

import (
	"math/rand"
)

type RandomizedSet struct {
	index map[int]int
	data  []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		index: make(map[int]int),
		data:  make([]int, 0),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.index[val]
	if ok {
		return false
	}

	this.data = append(this.data, val)
	this.index[val] = len(this.data) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.index[val]
	if !ok {
		return false
	}

	lastIdx := len(this.data) - 1
	if lastIdx != idx {
		this.data[idx] = this.data[lastIdx]
		this.index[this.data[idx]] = idx
	}
	this.data = this.data[:lastIdx]
	delete(this.index, val)

	return true
}

// -1
/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	l := len(this.data)
	if l == 0 {
		return l
	}
	return this.data[rand.Intn(l)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
