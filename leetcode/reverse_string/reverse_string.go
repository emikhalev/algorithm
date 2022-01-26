// https://leetcode.com/problems/reverse-string/
package reverse_string

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		t := s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = t
	}
}
