// https://leetcode.com/problems/reorder-list/
package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	l := listLen(head)
	addr := addresses(head, l)
	l--
	for i := 0; i < l/2; i++ {
		moveLastNode(addr, i)
	}
}

func printList(head *ListNode) {
	c := head
	cnt := 0
	for c != nil {
		fmt.Printf("%v ", c.Val)
		c = c.Next
		cnt++
	}
	fmt.Println()
}

func moveLastNode(a []*ListNode, after int) {
	i := len(a) - after - 1

	a[i-1].Next = nil
	a[after].Next = a[i]
	a[i].Next = a[after+1]
}

func getAddr(a []*ListNode, i int) *ListNode {
	if i < 0 || i >= len(a) {
		return nil
	}
	return a[i]
}

func addresses(root *ListNode, l int) []*ListNode {
	c := root
	a := make([]*ListNode, l)
	i := 0
	for c != nil {
		a[i] = c
		c = c.Next
		i++
	}
	return a
}

func listLen(root *ListNode) int {
	i := 0
	c := root
	for c != nil {
		i++
		c = c.Next
	}
	return i
}
