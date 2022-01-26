// https://leetcode.com/problems/delete-node-in-a-linked-list/
package delete_node_in_a_linked_list

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

func deleteNode(node *ListNode) {
	n := node
	var pn *ListNode
	for n.Next != nil {
		n.Val = n.Next.Val
		pn = n
		n = n.Next
	}
	pn.Next = nil
}
