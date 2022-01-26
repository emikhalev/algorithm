// https://leetcode.com/problems/divide-two-integers/
package divide_two_integers

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 0 {
		return 2<<31 - 1
	}
	i := 0
	sign := 1
	if divisor < 0 {
		divisor = -divisor
		sign = -sign
	}
	if dividend < 0 {
		dividend = -dividend
		sign = -sign
	}
	for dividend >= 0 {
		dividend -= divisor
		i++
	}
	res := sign * (i - 1)

	MAX_INT := 1<<31 - 1
	MIN_INT := -1 << 31
	if res > MAX_INT {
		res = MAX_INT
	}
	if res < MIN_INT {
		res = MIN_INT
	}
	return res
}
