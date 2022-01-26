// https://leetcode.com/problems/valid-palindrome/
package valid_palindrome

import (
	"strings"
)

func isPalindrome(s string) bool {
	bs := removeNonAlpa(s)

	middle := len(bs) / 2
	for i := 0; i < middle; i++ {
		if bs[i] != bs[len(bs)-i-1] {
			return false
		}
	}
	return true
}

func removeNonAlpa(s string) []byte {
	s = strings.ToLower(s)
	r := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if isAlpha(s[i]) {
			r = append(r, s[i])
		}
	}
	return r
}

func isAlpha(ch byte) bool {
	if ch >= 'a' && ch <= 'z' {
		return true
	}
	if ch >= 'A' && ch <= 'Z' {
		return true
	}
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}
