// https://leetcode.com/problems/happy-number/
package main

func isHappy(n int) bool {
	nums := make(map[int]struct{})
	for {
		nn := calc(n)
		if nn == 1 {
			return true
		}
		if nn <= 0 {
			return false
		}
		if _, ok := nums[nn]; ok {
			return false
		}
		nums[nn] = struct{}{}
		n = nn
	}
	return false
}

func calc(n int) int {
	r := int(0)
	for ; n > 0; n /= 10 {
		c := n % 10
		r = r + c*c
	}
	return r
}
