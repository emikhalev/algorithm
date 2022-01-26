// https://leetcode.com/problems/expression-add-operators/
package expression_add_operators

import (
	"strconv"
)

func addOperators(num string, target int) []string {
	res := make([]string, 0)
	helper([]byte(num), 0, 0, "", &res, 0, 0, target)
	return res
}

func helper(num []byte, pos int, preCalc int, prefix string, res *[]string,
	prevOperand, prevOp int, target int) {

	if pos >= len(num) && preCalc == target {
		(*res) = append((*res), prefix)
		return
	}

	a := preCalc
	i := pos
	for j := 0; i+j < len(num); j++ {
		bS := string(num[i : i+j+1])
		if len(bS) > 1 && bS[0] == '0' {
			continue
		}
		b, _ := strconv.Atoi(bS)
		bS = strconv.Itoa(b)

		if i > 0 {
			helper(num, i+j+1, a+b, prefix+"+"+bS, res, b, 1, target)
			helper(num, i+j+1, a-b, prefix+"-"+bS, res, b, 2, target)

			s := 0
			if prevOp == 1 { // +
				b = b * prevOperand
				s = (a - prevOperand) + b
			} else if prevOp == 2 { // -
				b = b * prevOperand
				s = (a + prevOperand) - b
			} else {
				prevOp = 3
				s = a * b
			}
			helper(num, i+j+1, s, prefix+"*"+bS, res, b, prevOp, target)
		} else {
			helper(num, i+j+1, b, bS, res, 0, 0, target)
		}
	}

}
