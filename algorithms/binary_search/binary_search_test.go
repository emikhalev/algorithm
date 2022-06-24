package binary_search

import (
	"testing"
)

type testCase struct {
	arr      []int
	target   int
	expIdx   int
	expFound bool
}

func TestBinarySearch(t *testing.T) {
	cases := []testCase{
		// Regular case
		{
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:   7,
			expIdx:   6,
			expFound: true,
		},
		// Empty array
		{
			arr:      []int{},
			target:   7,
			expIdx:   0,
			expFound: false,
		},
		// Single element array doesn't contains element
		{
			arr:      []int{1},
			target:   7,
			expIdx:   0,
			expFound: false,
		},
		// Single element array contains element
		{
			arr:      []int{7},
			target:   7,
			expIdx:   0,
			expFound: true,
		},
		// No value (in the middle)
		{
			arr:      []int{-1, -2, 2, 4, 6, 7, 9},
			target:   5,
			expIdx:   0,
			expFound: false,
		},
		// No value (less than min value in the array)
		{
			arr:      []int{-10, -2, -1, 4, 6, 7, 9},
			target:   -30,
			expIdx:   0,
			expFound: false,
		},
		// No value (more than max value in the array)
		{
			arr:      []int{-10, -2, -1, 4, 6, 7, 9},
			target:   50,
			expIdx:   0,
			expFound: false,
		},
		// First element
		{
			arr:      []int{-10, -2, -1, 4, 6, 7, 9},
			target:   -10,
			expIdx:   0,
			expFound: true,
		},
		// Last element
		{
			arr:      []int{-12, -10, -2, -1, 4, 6, 7, 9},
			target:   9,
			expIdx:   7,
			expFound: true,
		},
	}

	for caseNum, tc := range cases {
		idx, found := BinarySearch(tc.arr, tc.target)
		if found != tc.expFound {
			t.Errorf("case %d error: expecting found=%v, got: %v\n", caseNum, tc.expFound, found)
		}
		if idx != tc.expIdx {
			t.Errorf("case %d error: expecting idx=%d, got %d\n", caseNum, tc.expIdx, idx)
		}
	}
}
