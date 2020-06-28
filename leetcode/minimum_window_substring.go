// https://leetcode.com/problems/minimum-window-substring/
package main

func minWindow(s string, t string) string {
	ans := ""

	signT := signature(t)
	cntT := len(t)

	signS := make(map[byte]int)
	cntS := 0

	i, j := 0, 0
	for i < len(s) && j < len(s) {
		// skip all characters until first that can be found in signT
		for ; i < len(s); i++ {
			if _, ok := signT[s[i]]; ok {
				break
			}
		}
		if len(s) == i {
			break
		}
		if j < i {
			j = i
		}

		ch := s[j]
		j++

		cCntS, ok := signS[ch]
		if !ok {
			cCntS = 0
		}
		cCntS = cCntS + 1
		signS[ch] = cCntS

		if cCntT, ok := signT[ch]; ok && cCntS <= cCntT {
			cntS++
		}

		if cntS == cntT {
			// Move i ahead logic
			// -- find first substr with one character left and min substr
			for cntS == cntT {
				cst := string([]byte(s)[i:j])
				if len(ans) > len(cst) || len(ans) == 0 {
					ans = cst
				}

				ch := s[i]
				cCntS, _ := signS[ch]
				cCntS--
				signS[ch] = cCntS

				if cCntT, ok := signT[ch]; ok && cCntS < cCntT {
					cntS--
				}
				i++
			}
		}
	}

	return ans
}

func signature(st string) map[byte]int {
	r := make(map[byte]int)
	for i := 0; i < len(st); i++ {
		v, ok := r[st[i]]
		if !ok {
			v = 0
		}
		v++
		r[st[i]] = v
	}
	return r
}
