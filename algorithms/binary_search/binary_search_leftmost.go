package binary_search

func BinarySearchLeftmost(a []int, target int) (int, bool) {
	l, r := 0, len(a)
	for l < r {
		m := (l + r) / 2 // similar to floor((l+r)/2) https://go.dev/ref/spec#Arithmetic_operators, see Integer operators
		if a[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l, l < len(a) && a[l] == target
}
