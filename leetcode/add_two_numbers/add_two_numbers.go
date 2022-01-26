// https://leetcode.com/problems/add-two-numbers/
package add_two_numbers

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur1 := l1
	cur2 := l2
	r := new(ListNode)
	cur := r

	carry := 0
	for cur1 != nil || cur2 != nil {
		val1 := 0
		if cur1 != nil {
			val1 = cur1.Val
		}

		val2 := 0
		if cur2 != nil {
			val2 = cur2.Val
		}

		sum := carry + val1 + val2

		cur.Val = sum % 10
		carry = sum / 10

		if cur1 != nil {
			cur1 = cur1.Next
		}
		if cur2 != nil {
			cur2 = cur2.Next
		}

		if cur1 != nil || cur2 != nil {
			cur.Next = new(ListNode)
			cur = cur.Next
		}
	}

	if carry > 0 {
		cur.Next = new(ListNode)
		cur.Next.Val = carry
	}

	return r
}
