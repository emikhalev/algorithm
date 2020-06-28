// https://leetcode.com/problems/multiply-strings/
package main

func multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 {
		return ""
	}

	n1 := len(num1)
	n2 := len(num2)
	products := make([][]byte, n2)
	shift := 0
	for i := n2 - 1; i >= 0; i-- {
		products[i] = mulSingle(toBytes(num1), num2[i]-'0', n2+n1, shift)
		shift++
	}
	return toChars(add(products, n2+n1))
}

func toBytes(s string) []byte {
	n := []byte(s)
	for i, v := range n {
		n[i] = v - '0'
	}
	return n
}

func toChars(n []byte) string {
	for i, v := range n {
		n[i] = v + '0'
	}
	return string(n)
}

func add(nums [][]byte, size int) []byte {
	carry := int(0)
	res := make([]byte, size+1)
	rIdx := size
	for i := size - 1; i >= 0; i-- {
		s := carry
		for j := range nums {
			s = s + int(nums[j][i])
		}

		carry = s / 10

		v := byte(s % 10)
		res[rIdx] = v
		rIdx--
	}

	for ; carry != 0; carry /= 10 {
		res[rIdx] = byte(carry % 10)
		rIdx--
	}

	return trimLeft(res)
}

func trimLeft(r []byte) []byte {
	i := 0
	for ; i < len(r)-1 && r[i] == 0; i++ {
	}
	return r[i:]
}

func mulSingle(num1 []byte, a byte, size, shift int) []byte {
	carry := byte(0)
	rIdx := size - shift - 1
	res := make([]byte, size)
	for i := len(num1) - 1; i >= 0; i-- {
		b := num1[i]
		pd := a*b + carry

		carry = pd / 10
		res[rIdx] = pd % 10
		rIdx--
	}

	if carry != 0 {
		res[rIdx] = carry
		rIdx--
	}

	return res
}
