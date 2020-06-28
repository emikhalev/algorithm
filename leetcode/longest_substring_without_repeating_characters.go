// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	maxLen := 0
	i := 0
	for j := 0; j < len(s); j++ {
		ch := s[j]
		if pos, ok := m[ch]; ok {
			i = max(i, pos)
		}

		m[ch] = j + 1

		curLen := j - i + 1
		maxLen = max(maxLen, curLen)
	}
	return maxLen
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
