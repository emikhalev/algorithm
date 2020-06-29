// https://leetcode.com/problems/valid-palindrome-ii/
package main

func validPalindrome(s string) bool {
	mid := len(s) / 2
	for i := 0; i < mid; i++ {
		if s[i] != s[len(s)-i-1] {
			j := len(s) - i - 1
			if !isPolyndrome(s, i+1, j) && !isPolyndrome(s, i, j-1) {
				return false
			}
		}
	}
	return true
}

func isPolyndrome(s string, l, r int) bool {
	mid := l + (r-l)/2
	for i := l; i <= mid; i++ {
		if s[i] != s[r-i+l] {
			return false
		}
	}
	return true
}
