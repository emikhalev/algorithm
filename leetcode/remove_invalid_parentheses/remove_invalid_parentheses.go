// https://leetcode.com/problems/remove-invalid-parentheses/

package remove_invalid_parentheses

func removeInvalidParentheses(s string) []string {
	toRemove := parToRemove(s)
	r := processParenthesesAtPos(s, getNextParentheses(s, -1), toRemove)
	r = removeDuplicates(r)
	return r
}

func removeDuplicates(r []string) []string {
	res := make([]string, 0, len(r))
	m := make(map[string]struct{})
	for i := 0; i < len(r); i++ {
		if _, ok := m[r[i]]; ok {
			continue
		}
		res = append(res, r[i])
		m[r[i]] = struct{}{}
	}
	return res
}

func parToRemove(s string) int {
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		}
		if s[i] == ')' {
			if left > 0 {
				left--
			} else {
				right++
			}
		}
	}
	return left + right
}

func isValid(s string) bool {
	return parToRemove(s) == 0
}

func processParenthesesAtPos(s string, pos int, toRemove int) []string {
	if toRemove == 0 {
		if isValid(s) {
			return []string{s}
		}
		return nil
	}

	if pos <= -1 {
		return nil
	}

	nextPos := getNextParentheses(s, pos)
	s1 := processParenthesesAtPos(s, nextPos, toRemove)
	s2 := processParenthesesAtPos(removeChar(s, pos), nextPos-1, toRemove-1)

	return append(s1, s2...)
}

func removeChar(s string, pos int) string {
	if s == "" {
		return s
	}
	idx := 0
	bs := make([]byte, len(s)-1)
	for i := 0; i < len(s); i++ {
		if i == pos {
			continue
		}
		bs[idx] = s[i]
		idx++
	}
	return string(bs)
}

func getNextParentheses(s string, pos int) int {
	if s == "" {
		return -1
	}
	np := -1
	for i := pos + 1; i < len(s); i++ {
		if s[i] == ')' || s[i] == '(' {
			np = i
			break
		}
	}
	return np
}
