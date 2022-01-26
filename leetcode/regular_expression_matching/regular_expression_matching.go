// https://leetcode.com/problems/regular-expression-matching/
package regular_expression_matching

func isMatch(s string, p string) bool {
	return checkMatch(s, p, 0, 0)
}

func checkMatch(s string, re string, sPos, rePos int) bool {
	for rePos < len(re) && sPos < len(s) {

		ch := re[rePos]
		nCh := byte(0)
		if rePos < len(re)-1 {
			nCh = re[rePos+1]
		}

		if ch == '*' {
			return false
		}

		if ch == '.' {
			if nCh != '*' {
				sPos++
				rePos++
				continue
			} else {
				for i := sPos; i <= len(s); i++ {
					if checkMatch(s, re, i, rePos+2) {
						return true
					}
				}
				return false
			}
		}

		if isLetter(ch) {
			if nCh != '*' {
				if s[sPos] == ch {
					sPos++
					rePos++
					continue
				}
				return false
			} else {
				for i := sPos; i < len(s) && s[i] == ch; i++ {
					if checkMatch(s, re, i+1, rePos+2) {
						return true
					}
				}
				rePos += 2
				continue
			}
		}

	}

	for rePos < len(re) {
		nCh := byte(0)
		if rePos < len(re)-1 {
			nCh = re[rePos+1]
		}
		if nCh == '*' {
			rePos += 2
			continue
		}
		return false
	}

	return len(s) == sPos && rePos >= len(re)
}

func isLetter(a byte) bool {
	return a >= 'a' && a <= 'z'
}
