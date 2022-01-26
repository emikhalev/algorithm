// https://leetcode.com/problems/string-to-integer-atoi/
package string_to_integer_atoi

func myAtoi(str string) int {

	MAX_INT := 1<<31 - 1
	MIN_INT := -1 * 1 << 31

	if len(str) == 0 {
		return 0
	}

	i := 0
	for i = 0; i < len(str) && str[i] == 32; i++ {
	}
	if i >= len(str) {
		return 0
	}

	sign := 1
	if isSign(str[i]) {
		if str[i] == '-' {
			sign = -sign
		}
		i++
	}

	num := 0
	for i := i; i < len(str); i++ {
		ch := str[i]
		if !isDigit(ch) {
			break
		}

		v := ch - '0'
		if num == 0 {
			num = int(v)
			continue
		}
		num = num*10 + int(v)

		if num*sign > MAX_INT {
			return MAX_INT
		}
		if num*sign < MIN_INT {
			return MIN_INT
		}

	}

	return sign * num
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isSign(ch byte) bool {
	return ch == '-' || ch == '+'
}
