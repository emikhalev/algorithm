// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/
package minimum_remove_to_make_valid_parentheses

func minRemoveToMakeValid(s string) string {
	bs, _ := into([]byte(s), 0, 0)
	return string(bs)
}

func into(s []byte, pos, level int) ([]byte, int) {
	for i := pos; i < len(s); i++ {
		c := s[i]
		if c == '(' {
			s, i = into(s, i+1, level+1)
			continue
		}
		if c == ')' {
			if level > 0 {
				return s, i
			} else {
				if i+1 < len(s) {
					s = append(s[:i], s[i+1:]...)
				} else {
					s = s[:i]
				}
				i--
			}
		}
	}
	if level > 0 {
		if pos < len(s) {
			s = append(s[:pos-1], s[pos:]...)
		} else {
			s = s[:pos-1]
		}
		return s, len(s) - 1
	}
	return s, len(s)
}
