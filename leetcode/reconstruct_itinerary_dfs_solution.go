// https://leetcode.com/problems/reconstruct-itinerary/submissions/

package main

import (
	"sort"
)

func findItinerary(tickets [][]string) []string {
	r := NewReconstruct("JFK", tickets)
	return r.Itinerary()
}

type Reconstruct struct {
	adjList     map[string][]string
	edgeUsed    map[string][]bool
	used, total int
	path        []string
	startIATA   string
}

func NewReconstruct(startIATA string, tickets [][]string) *Reconstruct {
	r := new(Reconstruct)
	r.total = len(tickets)
	r.startIATA = startIATA
	r.makeAdjList(tickets)
	return r
}

func (r *Reconstruct) makeAdjList(tickets [][]string) {
	r.adjList = make(map[string][]string)
	for _, ticket := range tickets {
		dep, arr := ticket[0], ticket[1]
		adj := r.adjList[dep]
		r.adjList[dep] = append(adj, arr)
	}
	for iata, tickets := range r.adjList {
		sort.Slice(tickets, func(i, j int) bool { return tickets[i] < tickets[j] })
		r.adjList[iata] = tickets
	}
}

func (r *Reconstruct) Itinerary() []string {
	r.used = 0
	r.path = make([]string, 0)
	r.path = append(r.path, r.startIATA)
	r.edgeUsed = make(map[string][]bool)
	r.dfs(r.startIATA)
	return r.path
}

func (r *Reconstruct) dfs(iata string) {
	for arrIdx, arrIATA := range r.adjList[iata] {

		if r.isEdgeUsed(iata, arrIdx) {
			continue
		}

		r.used++
		r.path = append(r.path, arrIATA)
		r.setEdgeUsed(iata, arrIdx, true)

		r.dfs(arrIATA)
		if r.used == r.total {
			return
		}

		r.setEdgeUsed(iata, arrIdx, false)
		r.path = r.path[:len(r.path)-1]
		r.used--
	}
}

func (r *Reconstruct) setEdgeUsed(iata string, arrIdx int, isUsed bool) {
	usedEdges, ok := r.edgeUsed[iata]
	if !ok {
		usedEdges = make([]bool, len(r.adjList[iata]))
	}
	usedEdges[arrIdx] = isUsed
	r.edgeUsed[iata] = usedEdges
}

func (r *Reconstruct) isEdgeUsed(iata string, arrIdx int) bool {
	usedEdges, ok := r.edgeUsed[iata]
	if !ok {
		usedEdges = make([]bool, len(r.adjList[iata]))
	}
	return usedEdges[arrIdx]
}
