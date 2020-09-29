// https://leetcode.com/problems/zigzag-conversion/

package main

func convert(s string, numRows int) string {
	hopSize := (numRows - 1) + (numRows - 1)
	if hopSize == 0 {
		return s
	}

	numHops := len(s)/hopSize + 1

	res := make([]byte, 0, len(s))
	for r := 0; r < numRows; r++ {
		for j := 0; j < numHops; j++ {
			c := j * hopSize
			idx := r + c
			if idx >= len(s) {
				break
			}
			res = append(res, s[idx])

			sidx := (r + c) + (hopSize - r*2)
			if r > 0 && r < numRows-1 && sidx < len(s) {
				res = append(res, s[sidx])
			}
		}
	}

	return string(res)
}
