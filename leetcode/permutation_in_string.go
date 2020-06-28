// https://leetcode.com/problems/permutation-in-string/
package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	sig1 := make([]int, 26)
	sig2 := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		sig1[s1[i]-'a']++
		sig2[s2[i]-'a']++
	}

	for i := 0; i < len(s2)-len(s1); i++ {
		if sigEqual(sig1, sig2) {
			return true
		}

		sig2[s2[i]-'a']--
		sig2[s2[i+len(s1)]-'a']++
	}

	return sigEqual(sig1, sig2)
}

func sigEqual(sig1, sig2 []int) bool {
	for i := 0; i < len(sig1); i++ {
		if sig1[i] != sig2[i] {
			return false
		}
	}
	return true
}
