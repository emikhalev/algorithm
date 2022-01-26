package add_two_numbers_ii

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	cases := []struct {
		l1Param []int
		l2Param []int
		exp     []int
	}{
		{
			l1Param: []int{7, 2, 4, 3},
			l2Param: []int{5, 6, 4},
			exp:     []int{7, 8, 0, 7},
		},
		{
			l1Param: []int{2, 4, 3},
			l2Param: []int{5, 6, 4},
			exp:     []int{8, 0, 7},
		},
		{
			l1Param: []int{0},
			l2Param: []int{0},
			exp:     []int{0},
		},
	}

	for idx, tc := range cases {
		l1 := arrayToList(tc.l1Param)
		l2 := arrayToList(tc.l2Param)
		res := addTwoNumbers(l1, l2)
		resArr := listToArray(res)

		if reflect.DeepEqual(resArr, tc.exp) == false {
			t.Errorf("case %d failed: expecting %v, got %v", idx, tc.exp, resArr)
		}
	}
}

func arrayToList(arr []int) *ListNode {
	var root *ListNode
	curr := root
	for _, el := range arr {
		node := &ListNode{
			Val:  el,
			Next: nil,
		}
		if curr == nil {
			root = node
			curr = node
		} else {
			curr.Next = node
			curr = curr.Next
		}
	}
	return root
}

func listToArray(root *ListNode) []int {
	res := make([]int, 0, 4)
	for curr := root; curr != nil; curr = curr.Next {
		res = append(res, curr.Val)
	}
	return res
}
