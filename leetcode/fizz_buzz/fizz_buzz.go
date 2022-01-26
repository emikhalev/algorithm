// https://leetcode.com/problems/fizz-buzz/
package fizz_buzz

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := 1; i <= n; i++ {
		ans[i-1] = ""
		if i%3 == 0 {
			ans[i-1] = "Fizz"
		}
		if i%5 == 0 {
			ans[i-1] = ans[i-1] + "Buzz"
		}
		if i%5 != 0 && i%3 != 0 {
			ans[i-1] = strconv.Itoa(i)
		}
	}
	return ans
}
