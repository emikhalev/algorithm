// https://leetcode.com/problems/time-based-key-value-store
package time_based_key_value_store

type TimeMap struct {
	data map[string][]timeMapRecord
}

type timeMapRecord struct {
	ts  int
	val string
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][]timeMapRecord),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	d, ok := this.data[key]
	if !ok {
		d = make([]timeMapRecord, 0)
	}
	d = append(d, timeMapRecord{
		ts:  timestamp,
		val: value,
	})
	this.data[key] = d
}

func (this *TimeMap) Get(key string, timestamp int) string {
	d, ok := this.data[key]
	if !ok {
		return ""
	}
	// binary search
	l := 0
	r := len(d) - 1
	idx := -1
	for l <= r {
		m := (l + r) / 2
		if d[m].ts < timestamp {
			l = m + 1
			idx = m
			continue
		} else if d[m].ts > timestamp {
			r = r - 1
		} else {
			idx = m
			break
		}
	}
	if idx == -1 {
		return ""
	}
	return d[idx].val
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
