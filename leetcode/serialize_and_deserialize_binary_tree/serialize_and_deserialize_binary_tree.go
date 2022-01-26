// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/
package serialize_and_deserialize_binary_tree

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	s := this.serializeNode(root)
	st := "["
	for i := 0; i < len(s); i++ {
		if len(st) > 1 {
			st += ","
		}
		if s[i] == nil {
			st = st + "null"
			continue
		}
		st += fmt.Sprintf("%v", *s[i])
	}
	st += "]"
	return st
}

func (this *Codec) serializeNode(n *TreeNode) []*int {
	if n == nil {
		return []*int{nil}
	}

	sM := []*int{&n.Val}
	sL := this.serializeNode(n.Left)
	sR := this.serializeNode(n.Right)

	sRes := append(sM, sL...)
	sRes = append(sRes, sR...)

	return sRes
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	db := []byte(data)
	db = db[1 : len(db)-1]
	parts := strings.Split(string(db), ",")

	root, _ := this.deserializeNode(parts, 0)
	return root
}

func (this *Codec) deserializeNode(nodes []string, n int) (*TreeNode, int) {
	if n > len(nodes) || nodes[n] == "null" {
		return nil, n + 1
	}
	v, _ := strconv.Atoi(nodes[n])
	node := &TreeNode{
		Val: v,
	}
	n++

	node.Left, n = this.deserializeNode(nodes, n)
	node.Right, n = this.deserializeNode(nodes, n)

	return node, n
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
