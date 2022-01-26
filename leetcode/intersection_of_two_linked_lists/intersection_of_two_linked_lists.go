// https://leetcode.com/problems/intersection-of-two-linked-lists/
package intersection_of_two_linked_lists

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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pA, pB := headA, headB
	var lA, lB *ListNode = nil, nil

	for lA == nil || lB == nil || lA == lB {
		if pA == pB {
			return pA
		}
		if pA.Next == nil {
			if lA == nil {
				lA = pA
			}
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB.Next == nil {
			if lB == nil {
				lB = pB
			}
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	return nil
}
