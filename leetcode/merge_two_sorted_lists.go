// https://leetcode.com/problems/merge-two-sorted-lists/
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	l1cur, l2cur := l1, l2
	r := new(ListNode)
	c := r
	for l1cur != nil && l2cur != nil {
		if l1cur.Val <= l2cur.Val {
			c.Val = l1cur.Val
			l1cur = l1cur.Next
		} else {
			c.Val = l2cur.Val
			l2cur = l2cur.Next
		}
		c.Next = new(ListNode)
		c = c.Next
	}

	c = copyList(l1cur, c)
	copyList(l2cur, c)

	return r
}

func copyList(fromList, toList *ListNode) *ListNode {
	if fromList == nil || toList == nil {
		return toList
	}
	pc := fromList
	for fromList != nil {
		toList.Val = fromList.Val

		fromList = fromList.Next

		pc = toList
		toList.Next = new(ListNode)
		toList = toList.Next
	}
	pc.Next = nil
	return pc
}
