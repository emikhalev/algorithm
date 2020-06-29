// https://leetcode.com/problems/license-key-formatting/
package main

func licenseKeyFormatting(S string, K int) string {
	dCnt := 0
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == '-' {
			dCnt++
		}
	}
	L := len(S) - dCnt
	if L == 0 {
		return ""
	}
	gCnt := (L / K)
	if L-gCnt*K > 0 {
		gCnt++
	}
	ans := make([]byte, L+(gCnt-1))

	aIdx := len(ans) - 1
	cCnt := 0
	for i := len(S) - 1; i >= 0 && aIdx >= 0; i-- {
		if S[i] == '-' {
			continue
		}
		ans[aIdx] = toUpper(S[i])
		aIdx--
		cCnt++
		if cCnt == K && aIdx >= 0 {
			ans[aIdx] = '-'
			aIdx--
			cCnt = 0
		}
	}

	return string(ans)
}

func toUpper(ch byte) byte {
	if ch >= 'a' && ch <= 'z' {
		ch = ch - 'a' + 'A'
	}
	return ch
}
