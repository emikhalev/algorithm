// https://leetcode.com/problems/merge-k-sorted-lists/
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	var h, nn *ListNode = nil, findNextMin(lists)
	c := h
	for ;nn!=nil; {
		if nn != nil {
			if h==nil {
				h = nn
				c = nn
			} else {
				c.Next = nn
				c = c.Next
			}
		}
		nn = findNextMin(lists)
	}
	return h
}

func findNextMin(lists []*ListNode) *ListNode {
	minIdx := -1
	for i:=0;i<len(lists);i++ {
		if lists[i]!=nil && (
			minIdx == -1 ||
				(lists[minIdx]!=nil && lists[minIdx].Val > lists[i].Val)) {
			minIdx = i
		}
	}
	if minIdx == -1 {
		return nil
	}

	min := lists[minIdx]
	lists[minIdx] = lists[minIdx].Next
	return min
}