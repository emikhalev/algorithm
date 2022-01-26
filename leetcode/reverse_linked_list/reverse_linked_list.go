// https://leetcode.com/problems/reverse-linked-list/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	c := head
	var p *ListNode = nil
	for c != nil {
		tn := c.Next
		c.Next = p
		p = c
		c = tn
	}
	return p
}

func printList(head *ListNode) {
	c := head
	for c != nil {
		fmt.Printf("%d ", c.Val)
		c = c.Next
	}
	fmt.Println()
}

func listFromSlice(sl []int) *ListNode {
	var head *ListNode = nil
	c := head
	for i := 0; i < len(sl); i++ {
		n := &ListNode{
			Val:  sl[i],
			Next: nil,
		}
		if c == nil {
			c = n
			head = n
		}
		c.Next = n
		c = n
	}
	return head
}

func main() {
	ll := listFromSlice([]int{1, 2, 3, 4, 5})
	ll = reverseList(ll)
	printList(ll)
}
