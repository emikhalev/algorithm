// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	current := head
	cnt := 0
	for current != nil {
		current = current.Next
		cnt++
	}

	nPos := cnt - n - 1
	if nPos == -1 {
		return head.Next
	}
	if nPos < 0 {
		return nil
	}

	current = head
	for i := 0; i < nPos; i++ {
		current = current.Next
	}

	if current.Next != nil {
		current.Next = current.Next.Next
	}

	return head
}
