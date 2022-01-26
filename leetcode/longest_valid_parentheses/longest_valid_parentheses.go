// https://leetcode.com/problems/longest-valid-parentheses/
package longest_valid_parentheses

func longestValidParentheses(s string) int {

	maxL := 0

	// seq[number] = {}int{left, right}
	validSeqs := make([][]int, 0)
	l, r := 0, 1

	for {
		l, r = nextValidSeq(s, l)

		if l > r {
			break
		}

		l, r = wideValidSeq(s, l, r)
		validSeqs = append(validSeqs, []int{l, r})

		l = r + 1
		r = r + 1
	}

	merged := false
	wided := true
	for merged || wided {
		merged = false
		wided = false
		validSeqs, merged = mergeSeq(validSeqs)
		if merged {
			for i := 0; i < len(validSeqs); i++ {
				l, r := validSeqs[i][0], validSeqs[i][1]
				l, r = wideValidSeq(s, l, r)
				if l != validSeqs[i][0] || r != validSeqs[i][1] {
					wided = true
				}
				validSeqs[i][0], validSeqs[i][1] = l, r
			}
		}
	}

	maxL = findMaxSeq(validSeqs)

	return maxL
}

func mergeSeq(inSeq [][]int) (outSeq [][]int, merged bool) {
	if len(inSeq) == 0 {
		return inSeq, false
	}
	outSeq = make([][]int, 0, len(inSeq))

	l, r := inSeq[0][0], inSeq[0][1]
	merged = false
	for i := 1; i < len(inSeq); i++ {
		nl, nr := inSeq[i][0], inSeq[i][1]
		if (nl - r) <= 1 {
			r = nr
			merged = true
			continue
		}

		outSeq = append(outSeq, []int{l, r})

		l = nl
		r = nr
	}
	outSeq = append(outSeq, []int{l, r})

	return outSeq, merged
}

func nextValidSeq(s string, start int) (l, r int) {
	l, r = start, -1
	for ; l < len(s)-1; l++ {
		if s[l] == '(' && s[l+1] == ')' {
			r = l + 1
			break
		}
	}
	for i := r + 1; i < len(s)-1; i += 2 {
		if s[i] != '(' || s[i+1] != ')' {
			r = i - 1
			break
		}
		r = i + 1
	}
	return l, r
}

func wideValidSeq(s string, left, right int) (l, r int) {
	l = left
	r = right
	for i := 1; left-i >= 0 && right+i < len(s); i++ {
		if s[left-i] != '(' || s[right+i] != ')' {
			break
		}
		l = left - i
		r = right + i
	}
	return l, r
}

func findMaxSeq(validSeqs [][]int) int {
	maxCnt := -1
	for i := 0; i < len(validSeqs); i++ {
		l, r := validSeqs[i][0], validSeqs[i][1]
		if maxCnt < r-l {
			maxCnt = r - l
		}
	}
	return maxCnt + 1
}

func bruteForce(s string) int {
	for level := len(s); level >= 1; level-- {
		for i := 0; i < len(s)-(level-1); i++ {
			cS := s[i : i+level]
			if isValid(cS) {
				return len(cS)
			}
		}
	}
	return 0
}

func isValid(s string) bool {
	openCnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			openCnt++
		}
		if s[i] == ')' {
			openCnt--
		}
		if openCnt < 0 {
			return false
		}
	}
	return openCnt == 0
}
