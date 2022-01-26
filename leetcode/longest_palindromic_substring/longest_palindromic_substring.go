// https://leetcode.com/problems/longest-palindromic-substring/
package longest_palindromic_substring

func longestPalindrome(s string) string {
	cnt := 0
	str := ""
	for i := 0; i < len(s); i++ {
		cCnt, cStr := maxPalindrome(s, i, i)
		if cnt < cCnt {
			cnt = cCnt
			str = cStr
		}

		cCnt, cStr = maxPalindrome(s, i, i+1)
		if cnt < cCnt {
			cnt = cCnt
			str = cStr
		}
	}
	return str
}

func maxPalindrome(s string, i, j int) (int, string) {
	cnt := 0
	for i >= 0 && j < len(s) {
		if s[i] == s[j] {
			cnt++
			if i != j {
				cnt++
			}
		} else {
			break
		}
		i--
		j++
	}
	return cnt, string([]byte(s)[(i + 1):j])
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
