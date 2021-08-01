// https://leetcode.com/problems/reverse-nodes-in-k-group/
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	curr := head
	firstReverse := true
	var prev *ListNode = nil
	for curr != nil {
		last := curr // iter2: last = &3
		for i := 0; i < k-1; i++ {
			last = last.Next
			if last == nil {
				return head
			}
		}
		nextCurr := last.Next

		reverseInPlace(curr, k)
		if prev != nil {
			prev.Next = last
		}

		if firstReverse {
			head = last
			firstReverse = false
		}
		prev = curr
		curr = nextCurr
	}
	return head
}

func reverseInPlace(root *ListNode, k int) {
	curr := root
	next := curr.Next
	for i := 0; i < k-1; i++ {
		tmp := next.Next
		next.Next = curr
		curr = next
		next = tmp
	}
	root.Next = next
}
