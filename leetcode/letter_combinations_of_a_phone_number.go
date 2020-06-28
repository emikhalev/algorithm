// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
package main

func letterCombinations(digits string) []string {
	letters := make([][]byte, 0, len(digits))
	for i := 0; i < len(digits); i++ {
		d := digits[i] - '0'
		ls := number2Letters(d)
		letters = append(letters, ls)
	}

	combinations := make([]string, 0)
	combinations = getCombinations("", letters, combinations, 0)

	return combinations
}

func getCombinations(prefix string, letters [][]byte, combinations []string, level int) []string {
	if level >= len(letters) {
		return combinations
	}
	for i := 0; i < len(letters[level]); i++ {
		prefix := prefix + string(letters[level][i])
		if level < len(letters)-1 {
			combinations = getCombinations(prefix, letters, combinations, level+1)
			continue
		}
		combinations = append(combinations, prefix)
	}
	return combinations
}

func number2Letters(number byte) []byte {
	if number < 0 || number > 9 {
		return []byte{}
	}
	n2l := [][]byte{
		0: {' '},
		1: {},
		2: {'a', 'b', 'c'},
		3: {'d', 'e', 'f'},
		4: {'g', 'h', 'i'},
		5: {'j', 'k', 'l'},
		6: {'m', 'n', 'o'},
		7: {'p', 'q', 'r', 's'},
		8: {'t', 'u', 'v'},
		9: {'w', 'x', 'y', 'z'},
	}
	return n2l[number]
}
