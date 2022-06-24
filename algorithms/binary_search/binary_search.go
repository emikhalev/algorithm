package binary_search

func BinarySearch(a []int, target int) (int, bool) {
	l, r := 0, len(a)-1
	for l <= r {
		m := (l + r) / 2
		if a[m] > target {
			r = m - 1
		} else if a[m] < target {
			l = m + 1
		} else {
			return m, true
		}
	}
	return 0, false
}
