// https://leetcode.com/problems/find-duplicate-file-in-system/
package main

import (
	"strings"
)

func findDuplicate(paths []string) [][]string {
	files := make(map[string][]string) // contentHash -> []paths
	for _, pathInfo := range paths {
		parts := strings.Split(pathInfo, " ")
		if len(parts) <= 1 {
			continue
		}

		dir := parts[0]
		for k := 1; k < len(parts); k++ {
			name, cont := fileNameContent(parts[k])
			files[cont] = append(files[cont], dir+"/"+name)
		}
	}

	ans := make([][]string, 0, len(files))
	i := 0
	for _, v := range files {
		if len(v) <= 1 {
			continue
		}
		ans = append(ans, v)
		i++
	}
	return ans
}

func fileNameContent(s string) (name, content string) {
	name = ""
	content = ""
	isCont := false
	for i := 0; i < len(s) && s[i] != ')'; i++ {
		if s[i] == '(' {
			isCont = true
			continue
		}
		if isCont {
			content += string(s[i])
		} else {
			name += string(s[i])
		}
	}
	return name, content
}
