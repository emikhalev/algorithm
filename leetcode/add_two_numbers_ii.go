// https://leetcode.com/problems/add-two-numbers-ii/
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ll1 := LLen(l1)
	ll2 := LLen(l2)
	var (
		h     *ListNode = nil
		carry int       = 0
	)

	if ll1 < ll2 {
		h, carry = sum(l2, l1, (ll2 - ll1), 0)
	} else {
		h, carry = sum(l1, l2, (ll1 - ll2), 0)
	}

	if carry > 0 {
		n := &ListNode{
			Val:  carry % 10,
			Next: h,
		}
		h = n
	}
	return h
}

func LLen(h *ListNode) int {
	c := 0
	for ; h != nil; h = h.Next {
		c++
	}
	return c
}

func sum(cl1, cl2 *ListNode, shift int, lvl int) (n *ListNode, carry int) {
	if cl1 == nil && cl2 == nil {
		return nil, 0
	}

	nl1 := cl1.Next
	nl2 := cl2
	if lvl >= shift {
		nl2 = nl2.Next
	}
	nn, carry := sum(nl1, nl2, shift, lvl+1)

	v1 := cl1.Val
	v2 := 0
	if shift <= lvl {
		v2 = cl2.Val
	}
	s := carry + v1 + v2
	n = &ListNode{
		Val:  s % 10,
		Next: nn,
	}
	return n, s / 10
}
