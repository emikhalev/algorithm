// https://leetcode.com/problems/design-hashmap/
package main

type MyHashMap struct {
	dataSize int
	data     [][]*myHashMapValue
}

type myHashMapValue struct {
	key     int
	val     int
	deleted bool
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		data:     make([][]*myHashMapValue, 2),
		dataSize: 2,
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	keyIdx := key % this.dataSize
	if len(this.data[keyIdx]) == 0 {
		this.data[keyIdx] = make([]*myHashMapValue, 0, 16)
	}

	// replace
	for _, v := range this.data[keyIdx] {
		if v.deleted == false && v.key == key {
			v.val = value
			return
		}
	}

	// add if have free space
	for _, v := range this.data[keyIdx] {
		if v.deleted {
			v.key = key
			v.val = value
			v.deleted = false
			return
		}
	}
	// append if not have free space
	this.data[keyIdx] = append(this.data[keyIdx], &myHashMapValue{
		key:     key,
		val:     value,
		deleted: false,
	})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	keyIdx := key % this.dataSize
	for _, v := range this.data[keyIdx] {
		if v.deleted == false && v.key == key {
			return v.val
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	keyIdx := key % this.dataSize
	for _, v := range this.data[keyIdx] {
		if v.deleted == false && v.key == key {
			v.deleted = true
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
