// https://leetcode.com/problems/integer-to-english-words/
package main

import (
	"strings"
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	bil := (num / 1000000000)
	mil := (num / 1000000) % 1000
	th := (num / 1000) % 1000
	hn := (num % 1000)
	return strings.Trim(billionsToWord(bil) + millionsToWord(mil) + thousandsToWord(th) + hundredsToWord(hn)," ")
}

func billionsToWord(num int) string {
	if num == 0 {
		return ""
	}
	suf := ""
	if num>3 {
		suf = ""
	}
	ans := hundredsToWord(num)
	return ans + " Billion"+suf+" "
}

func millionsToWord(num int) string {
	if num == 0 {
		return ""
	}
	suf := ""
	if num>3 {
		suf = ""
	}
	ans := hundredsToWord(num)
	return ans + " Million"+suf+" "
}

func thousandsToWord(num int) string {
	if num == 0 {
		return ""
	}
	suf := ""
	if num>3 {
		suf = ""
	}
	ans := hundredsToWord(num)
	return ans + " Thousand"+suf+" "
}

func hundredsToWord(num int) string {
	if num == 0 {
		return ""
	}
	suf := ""
	if num/100>3 {
		suf = ""
	}
	ans := ""
	if num/100>0 {
		ans = singleToWord(num/100) + " Hundred"+suf+" "
	}
	ans = ans + decToWord(num%100)
	return strings.Trim(ans, " ")
}

func singleToWord(num int) string {
	switch num {
	case 1: return "One"
	case 2: return "Two"
	case 3: return "Three"
	case 4: return "Four"
	case 5: return "Five"
	case 6: return "Six"
	case 7: return "Seven"
	case 8: return "Eight"
	case 9: return "Nine"
	}
	return ""
}

func decToWord(num int) string {
	if num == 0 {
		return ""
	}
	if num<10 {
		return singleToWord(num)
	}
	if num>=10 && num<=19 {
		switch num {
		case 10: return "Ten"
		case 11: return "Eleven"
		case 12: return "Twelve"
		case 13: return "Thirteen"
		case 14: return "Fourteen"
		case 15: return "Fifteen"
		case 16: return "Sixteen"
		case 17: return "Seventeen"
		case 18: return "Eighteen"
		case 19: return "Nineteen"
		}
	}
	if num >= 20 && num<=99 {
		return strings.Trim(decToPrefix(num/10)+"ty "+singleToWord(num%10), " ")
	}
	return ""
}

func decToPrefix(num int) string  {
	v := singleToWord(num)
	switch num {
	case 2: return "Twen"
	case 3: return "Thir"
	case 4: return "For"
	case 5: return "Fif"
	case 8: return "Eigh"
	}
	return v
}