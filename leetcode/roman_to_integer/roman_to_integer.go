// https://leetcode.com/problems/roman-to-integer/
package roman_to_integer

func romanToInt(s string) int {
	v := 0
	i := 0
	for i = 0; i < len(s)-1; i++ {
		num := romanToNum(s[i])
		nextNum := romanToNum(s[i+1])

		if num < nextNum {
			v = v + nextNum - num
			i++
		} else {
			v = v + num
		}
	}

	if i == len(s)-1 {
		v = v + romanToNum(s[i])
	}

	return v
}

func romanToNum(num byte) int {
	switch num {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}
