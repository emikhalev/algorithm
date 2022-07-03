package binary_search

import "testing"

func TestBinarySearchLeftmost(t *testing.T) {
	cases := []testCase{
		// 0. Regular case
		{
			arr:      []int{3, 4, 5, 6, 7, 7, 8, 9},
			target:   7,
			expIdx:   4,
			expFound: true,
		},
		// 1. Empty array
		{
			arr:      []int{},
			target:   7,
			expIdx:   0,
			expFound: false,
		},
		// 2. Single element array doesn't contains element
		{
			arr:      []int{1},
			target:   7,
			expIdx:   0,
			expFound: false,
		},
		// 3. Single element array contains element
		{
			arr:      []int{7},
			target:   7,
			expIdx:   0,
			expFound: true,
		},
		// 4. No value (in the middle)
		{
			arr:      []int{-1, -2, 2, 4, 6, 7, 9},
			target:   5,
			expIdx:   0,
			expFound: false,
		},
		// 5. No value (less than min value in the array)
		{
			arr:      []int{-10, -2, -1, 4, 6, 7, 9},
			target:   -30,
			expIdx:   0,
			expFound: false,
		},
		// 6. No value (more than max value in the array)
		{
			arr:      []int{-10, -2, -1, 4, 6, 7, 9},
			target:   50,
			expIdx:   0,
			expFound: false,
		},
		// 7. First element
		{
			arr:      []int{-10, -2, -1, 4, 6, 7, 9},
			target:   -10,
			expIdx:   0,
			expFound: true,
		},
		// 8. Last element
		{
			arr:      []int{-12, -10, -2, -1, 4, 6, 7, 9},
			target:   9,
			expIdx:   7,
			expFound: true,
		},
	}

	for caseNum, tc := range cases {
		idx, found := BinarySearchLeftmost(tc.arr, tc.target)
		if found != tc.expFound {
			t.Errorf("case %d error: expecting found=%v, got: %v\n", caseNum, tc.expFound, found)
		}
		if !tc.expFound { // idx can be random in such case
			continue
		}
		if idx != tc.expIdx {
			t.Errorf("case %d error: expecting idx=%d, got %d\n", caseNum, tc.expIdx, idx)
		}
	}
}
