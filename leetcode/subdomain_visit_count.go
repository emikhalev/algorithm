// https://leetcode.com/problems/subdomain-visit-count
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	vMap := make(map[string]int)

	for _, visit := range cpdomains {
		parts := strings.Split(visit, " ")
		cnt, _ := strconv.Atoi(parts[0])
		domain := parts[1]
		idx := 0
		for idx != -1 {
			addVisits(vMap, domain, cnt)
			idx = strings.Index(domain, ".")
			domain = string([]byte(domain)[idx+1:])
		}
	}

	ans := make([]string, len(vMap))
	i := 0
	for domain, cnt := range vMap {
		ans[i] = fmt.Sprintf("%d %s", cnt, domain)
		i++
	}
	return ans
}

func addVisits(vMap map[string]int, domain string, cnt int) {
	cCnt, _ := vMap[domain]
	cCnt += cnt
	vMap[domain] = cCnt
}
