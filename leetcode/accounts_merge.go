// https://leetcode.com/problems/accounts-merge/
package main

import (
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	links := linkMap(accounts)
	return mergeDFS(accounts, links)
}

// maps email to nodes
func linkMap(accounts [][]string) map[string][]int {
	links := make(map[string][]int) // email -> [id1,id2,.., idN]

	for i := 0; i < len(accounts); i++ {
		addedLinks := make(map[string]struct{})
		for j := 1; j < len(accounts[i]); j++ { // From 1st, because 0 is the FirstName
			email := accounts[i][j]
			if v, ok := links[email]; !ok {
				addedLinks[email] = struct{}{}
				links[email] = []int{i}
			} else if _, added := addedLinks[email]; !added {
				v = append(v, i)
				links[email] = v
				addedLinks[email] = struct{}{}
			}
		}
	}
	return links
}

func mergeDFS(accounts [][]string, links map[string][]int) [][]string {
	res := make([][]string, 0, 16)
	visited := make(map[int]struct{})

	for i := 0; i < len(accounts); i++ {
		if _, ok := visited[i]; !ok {
			title := accounts[i][0]
			res = append(res, []string{title})
			resIdx := len(res) - 1

			merged := mergeDFSNode(accounts, links, visited, i)
			res[resIdx] = append(res[resIdx], merged...)
		}
	}

	// remove duplicates
	for i := 0; i < len(res); i++ {
		added := make(map[string]struct{})
		emails := []string{res[i][0]}
		for j := 1; j < len(res[i]); j++ {
			email := res[i][j]
			if _, ok := added[email]; !ok {
				added[email] = struct{}{}
				emails = append(emails, email)
			}
		}
		res[i] = emails
	}

	// sort
	for _, acc := range res {
		sort.Sort(sort.StringSlice(acc))
	}
	return res
}

func mergeDFSNode(accounts [][]string, links map[string][]int, visited map[int]struct{},
	i int) []string {

	visited[i] = struct{}{}

	res := make([]string, 0, len(accounts[i])-1)
	addRes := make([]string, 0, len(accounts[i])-1)

	// iterate over node emaills and DFS by graph edges
	for eIdx := 1; eIdx < len(accounts[i]); eIdx++ {
		email := accounts[i][eIdx]
		emailLinks, _ := links[email]
		for _, lnkIdx := range emailLinks {
			if lnkIdx == i {
				res = append(res, email)
			}

			if _, ok := visited[lnkIdx]; ok {
				continue
			}

			merged := mergeDFSNode(accounts, links, visited, lnkIdx)
			res = append(res, merged...)
		}
	}

	return append(res, addRes...)
}
