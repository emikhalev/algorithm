// https://leetcode.com/problems/reverse-nodes-in-k-group/
package solution_1

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	first := head
	var reversed, reversedLast *ListNode

	for first != nil {

		// [1,2,3,4,5], k=2; i=0..2 (0,1,2) last = &1, 0(last=&1), 1(last=&2), 2(last=&3),
		last := first
		for i := 0; i < k-1; i++ {
			last = last.Next
			if last == nil {
				if reversed == nil {
					reversed = first
				} else {
					reversedLast.Next = first
				}
				return reversed
			}
		}

		// [1,2,3,4,5], [2,1], reversed = lst [2,1], reversedLast = lstLast [2, 1, ]
		lst, lstLast := reverse(first, k) // returns reversed copy

		if reversed == nil {
			reversed = lst
			reversedLast = lstLast
		} else {
			reversedLast.Next = lst
			reversedLast = lstLast
		}

		first = last.Next // first = &3
	}
	return reversed
}

// reverse: [1, 2, 3, 4, 5], reversed = nil, current = &1, prev = &1
func reverse(root *ListNode, k int) (reversed, reversedLast *ListNode) {
	current := root
	prev := reversed

	for idx := 0; idx < k; idx++ {
		elem := &ListNode{current.Val, nil}
		elem.Next = prev       // 1
		reversed = elem        // 2
		prev = elem            // 3
		current = current.Next // 4
		if idx == 0 {
			reversedLast = elem
		}
	}

	return reversed, reversedLast
}
