// https://leetcode.com/problems/verifying-an-alien-dictionary/
package verifying_an_alien_dictionary

func isAlienSorted(words []string, order string) bool {
	os := orders(order)

	for i := 0; i < len(words)-1; i++ {
		if !isAsc(words[i], words[i+1], os) {
			return false
		}
	}

	return true
}

func isAsc(word1, word2 string, order map[byte]int) bool {
	n := max(len(word1), len(word2))
	for i := 0; i < n; i++ {
		c1 := byte(0)
		if i < len(word1) {
			c1 = word1[i]
		}
		c2 := byte(0)
		if i < len(word2) {
			c2 = word2[i]
		}

		o1 := getOrder(order, int(c1))
		o2 := getOrder(order, int(c2))

		if o1 < o2 {
			return true
		}

		if o1 > o2 {
			return false
		}
	}
	return true
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func getOrder(orders map[byte]int, v int) int {
	order, ok := orders[byte(v)]
	if !ok {
		return -1
	}
	return order
}

func orders(order string) map[byte]int {
	r := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		r[order[i]] = i
	}
	return r
}
