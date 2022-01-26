// https://leetcode.com/problems/reverse-words-in-a-string/
package reverse_words_in_a_string

func reverseWords(s string) string {
	words := make([]string, 0, 16)
	inW := false
	w := ""
	for i := 0; i < len(s); i++ {

		if !inW && s[i] != 32 {
			inW = true
			w = ""
		}
		if inW && s[i] == 32 {
			inW = false
			words = append(words, w)
		}

		if inW {
			w = w + string(s[i])
		}
	}
	if inW {
		words = append(words, w)
	}

	ans := ""
	for i := len(words) - 1; i >= 0; i-- {
		ans += words[i]
		if i > 0 {
			ans += " "
		}
	}

	return ans
}
