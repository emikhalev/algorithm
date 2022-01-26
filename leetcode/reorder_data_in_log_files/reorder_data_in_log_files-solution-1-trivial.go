// https://leetcode.com/problems/reorder-data-in-log-files/

package reorder_data_in_log_files

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {

	words := make([]string, 0, len(logs))
	digits := make([]string, 0, 16)

	for _, lw := range logs {
		if isDigit(rec(lw)[0]) {
			digits = append(digits, lw)
		} else {
			words = append(words, lw)
		}
	}

	sort.Slice(words, func(i, j int) bool {
		recI := rec(words[i])
		recJ := rec(words[j])
		if recI == recJ {
			return words[i] < words[j]
		}
		return recI < recJ
	})

	return append(words, digits...)
}

func rec(log string) string {
	return string([]byte(log)[strings.Index(log, " ")+1:])
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

/*
["8 fj dzz k", "5r 446 6 3", "63 gu psub", "5 ba iedrz", "6s 87979 5", "3r 587 01", "jc 3480612", "bb wsrd kp", "b aq cojj", "r5 6316 71"]
["a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo","a2 act car"]
["1 n u", "r 527", "j 893", "6 14", "6 82"]
["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]
[]
["1 n u"]
["j 893", "6 14"]
*/
