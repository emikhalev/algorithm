package add_binary

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
			[]string{"11", "1"},
			"100",
		},
		{
			[]string{"1010", "1011"},
			"10101",
		},
	}

	for idx, tc := range cases {
		out := addBinary(tc.params[0], tc.params[1])
		if !reflect.DeepEqual(out, tc.exp) {
			t.Errorf("case %d error: got: %v, expected: %v", idx, out, tc.exp)
		}
	}
}
