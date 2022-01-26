// https://leetcode.com/problems/find-all-anagrams-in-a-string/
package find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	ans := make([]int, 0)

	pSig, sSig := [26]byte{}, [26]byte{}
	for i := 0; i < len(p); i++ {
		pSig[p[i]-'a']++
		sSig[s[i]-'a']++
	}

	i := len(p)
	for ; i < len(s); i++ {
		if checkSig(sSig, pSig) {
			ans = append(ans, i-len(p))
		}

		sSig[s[i-len(p)]-'a']--
		sSig[s[i]-'a']++
	}
	if checkSig(sSig, pSig) {
		ans = append(ans, i-len(p))
	}

	return ans
}

func checkSig(s1, s2 [26]byte) bool {
	for i := 0; i < 26; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
