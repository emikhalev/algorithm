package add_strings

import (
	"reflect"
	"testing"
)

func Test_addBinary(t *testing.T) {
	cases := []struct {
		params []string
		exp    string
	}{
		{
			[]string{"11", "123"},
			"134",
		},
		{
			[]string{"456", "77"},
			"533",
		},
		{
			[]string{"0", "0"},
			"0",
		},
	}

	for idx, tc := range cases {
		out := addStrings(tc.params[0], tc.params[1])
		if !reflect.DeepEqual(out, tc.exp) {
			t.Errorf("case %d error: got: %v, expected: %v", idx, out, tc.exp)
		}
	}
}
