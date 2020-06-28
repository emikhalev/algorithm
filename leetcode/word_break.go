// https://leetcode.com/problems/word-break/
package main

func wordBreak(s string, wordDict []string) bool {
	if s == "" {
		return false
	}

	// Generate dict map
	dict := make(map[string]struct{})
	breakDict := make(map[int]bool)
	for i := 0; i < len(wordDict); i++ {
		dict[wordDict[i]] = struct{}{}
	}

	return canBreak(s, 0, dict, breakDict)
}

func canBreak(s string, pos int, dict map[string]struct{}, breakDict map[int]bool) bool {
	if pos == len(s) {
		return true
	}

	if v, ok := breakDict[pos]; ok {
		return v
	}

	cw := ""
	for i := pos; i < len(s); i++ {
		cw = cw + string(s[i])
		if _, ok := dict[cw]; ok {
			res := canBreak(s, i+1, dict, breakDict)
			if res {
				breakDict[pos] = true
				return true
			}
		}
	}

	breakDict[pos] = false
	return false
}
