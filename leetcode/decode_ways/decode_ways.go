// https://leetcode.com/problems/decode-ways/
package decode_ways

func numDecodings(s string) int {
	return decode([]byte(s), 0)
}

func decode(s []byte, pos int) int {
	if pos >= len(s) {
		return 1
	}

	cnt := 0

	d := int(s[pos] - '0')
	if isNum(d) {
		cnt += decode(s, pos+1)
	}

	if isNum(d) && pos < len(s)-1 {
		if isNum(d*10 + int(s[pos+1]-'0')) {
			cnt += decode(s, pos+2)
		}
	}

	return cnt
}

func isNum(num int) bool {
	if num < 1 || num > 26 {
		return false
	}
	return true
}
