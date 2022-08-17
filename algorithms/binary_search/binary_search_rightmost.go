package binary_search

func BinarySearchRightmost(a []int, target int) (int, bool) {
	l, r := 0, len(a)
	for l < r {
		m := (l + r) / 2 // similar to floor((l+r)/2) https://go.dev/ref/spec#Arithmetic_operators, see Integer operators
		if a[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	return r - 1, r > 0 && a[r-1] == target
}
