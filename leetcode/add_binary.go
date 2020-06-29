// https://leetcode.com/problems/add-binary/
package main

func addBinary(a string, b string) string {
	aIdxShift := 0
	bIdxShift := 0
	maxSize := 0
	if len(a) < len(b) {
		aIdxShift = len(b) - len(a)
		maxSize = len(b)
	} else {
		bIdxShift = len(a) - len(b)
		maxSize = len(a)
	}

	res := make([]byte, maxSize+2)
	add := 0
	i := maxSize - 1
	rIdx := len(res) - 1
	for ; i >= 0; i-- {
		a := getNum(a, i-aIdxShift, 0)
		b := getNum(b, i-bIdxShift, 0)
		sum := add + a + b

		res[rIdx] = byte(sum%2) + '0'
		rIdx--
		add = sum / 2
	}
	for add > 0 {
		res[rIdx] = byte(add%2) + '0'
		add = add / 2
		rIdx--
	}

	return string(res[rIdx+1:])
}

func getNum(s string, idx int, defVal int) int {
	if idx < 0 || idx > len(s) {
		return defVal
	}
	return int(s[idx] - '0')
}
