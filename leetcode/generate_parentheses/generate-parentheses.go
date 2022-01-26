// https://leetcode.com/problems/generate-parentheses
package generate_parentheses

func generateParenthesis(n int) []string {
	ans := make([]string, 0, 0)
	helper(n, n, true, "", &ans)
	return ans
}

func helper(opened, closed int, open bool, cSt string, res *[]string) {
	if open {
		cSt += "("
		opened--
	} else {
		cSt += ")"
		closed--
	}

	if opened > 0 {
		helper(opened, closed, true, cSt, res)
	}
	if closed > 0 && closed > opened {
		helper(opened, closed, false, cSt, res)
	}
	if opened == 0 && closed == 0 {
		(*res) = append((*res), cSt)
	}
}
