// https://leetcode.com/problems/validate-ip-address/
package main

import (
	"strings"
)

func validIPAddress(IP string) string {
	if isIPv4(IP) {
		nums := strings.Split(IP, ".")
		for _, v := range nums {
			if checkIPv4Num(v) != true {
				return "Neither"
			}
		}
		return "IPv4"
	} else if isIPv6(IP) {
		nums := strings.Split(IP, ":")
		for _, v := range nums {
			if checkIPv6Num(v) != true {
				return "Neither"
			}
		}
		return "IPv6"
	} else {
		return "Neither"
	}
}

func checkIPv4Num(s string) bool {
	if len(s) > 3 || s == "" {
		return false
	}

	num := 0
	j := 1
	for i := len(s) - 1; i >= 0; i-- {
		if !isDigit(s[i]) {
			return false
		}
		digit := int(s[i] - '0')
		if digit == 0 && i == 0 && len(s) > 1 {
			return false
		}
		num = num + j*digit
		j = j * 10
	}

	if num < 0 || num > 255 {
		return false
	}

	return true
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func checkIPv6Num(s string) bool {
	if len(s) > 4 || s == "" {
		return false
	}

	num := 0
	j := 1
	for i := len(s) - 1; i >= 0; i-- {
		if !isHex(s[i]) {
			return false
		}
		digit := hexToDec(s[i])
		if digit < 0 {
			return false
		}
		num = num + j*digit
		j = j * 16
	}

	if num < 0 || num > 1<<32-1 {
		return false
	}

	return true
}

func isHex(ch byte) bool {
	return (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F') || isDigit(ch)
}

func hexToDec(ch byte) int {
	num := -1

	if ch >= 'a' && ch <= 'f' {
		num = int(ch - 'a')
	} else if ch >= 'A' && ch <= 'F' {
		num = int(ch - 'A')
	} else if ch >= '0' && ch <= '9' {
		num = int(ch - '0')
	}

	return num
}

func isIPv4(s string) bool {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			cnt++
		}
	}
	return cnt == 3
}

func isIPv6(s string) bool {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ':' {
			cnt++
		}
	}
	return cnt == 7
}
