// https://leetcode.com/problems/powx-n/
package powx_n

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if x == 1 {
		return 1
	}

	if x == -1 {
		if n%2 == 0 {
			return 1
		}
		return -1
	}

	sign := 1
	if n < 0 {
		sign = -1
		n = -n
	}

	r := float64(1)
	curProduct := x
	for cN := 1; cN <= n; {
		if (cN & n) > 0 {
			r = r * curProduct
		}

		cN = cN << 1
		curProduct = curProduct * curProduct
	}

	if sign < 0 {
		r = 1 / r
	}

	return r
}
