// https://leetcode.com/problems/add-strings/
package add_strings

func addStrings(num1 string, num2 string) string {
	total := make([]byte, max(len(num1), len(num2))+1)

	i, j, k := len(num1)-1, len(num2)-1, len(total)-1
	aV, bV, carry := byte(0), byte(0), byte(0)
	for i >= 0 || j >= 0 {
		if i >= 0 {
			aV = num1[i] - '0'
		} else {
			aV = 0
		}
		if j >= 0 {
			bV = num2[j] - '0'
		} else {
			bV = 0
		}

		sum := aV + bV + carry
		total[k] = sum%10 + '0'
		carry = sum / 10

		i--
		j--
		k--
	}

	if carry > 0 {
		total[k] = carry%10 + '0'
		k--
	}

	return string(total[k+1:])
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
