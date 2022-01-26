// https://leetcode.com/problems/lru-cache/
package lru_cache

type CacheNode struct {
	key int
	val int
	ts  int64
}

type LRUCache struct {
	data         []*CacheNode
	dataIdxByKey map[int]int
	idx          int
}

var (
	curTime int64 = 0
)

func Constructor(capacity int) LRUCache {
	curTime++

	return LRUCache{
		make([]*CacheNode, capacity),
		make(map[int]int),
		0,
	}
}

func (this *LRUCache) Get(key int) int {
	curTime++

	if this.idx == 0 {
		return -1
	}

	i := this.getKeyIdx(key)
	if i == -1 {
		return -1
	}

	this.data[i].ts = curTime
	return this.data[i].val
}

func (this *LRUCache) Put(key int, value int) {
	curTime++

	freeIdx := this.getKeyIdx(key)

	if freeIdx == -1 {
		// If cache is full - find LRU element
		if this.idx == len(this.data) {
			minTime := int64(1<<63 - 1)
			for i := 0; i < len(this.data); i++ {
				if minTime > this.data[i].ts {
					freeIdx = i
					minTime = this.data[i].ts
				}
			}
			if freeIdx != -1 {
				delete(this.dataIdxByKey, this.data[freeIdx].key)
			}
		} else {
			freeIdx = this.idx
			this.idx++
		}
	}

	// Add
	this.data[freeIdx] = &CacheNode{
		key: key,
		val: value,
		ts:  curTime,
	}
	this.dataIdxByKey[key] = freeIdx
}

func (this *LRUCache) getKeyIdx(key int) int {
	if i, ok := this.dataIdxByKey[key]; ok {
		return i
	}
	return -1
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
