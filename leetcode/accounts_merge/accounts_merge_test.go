package accounts_merge

import (
	"reflect"
	"testing"
)

func Test_accountsMerge(t *testing.T) {
	cases := []struct {
		params [][]string
		exp    [][]string
	}{
		{
			[][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}},
			[][]string{{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}},
		},
		{
			[][]string{{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"}, {"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"}, {"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"}, {"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"}, {"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"}},
			[][]string{{"Ethan", "Ethan0@m.co", "Ethan4@m.co", "Ethan5@m.co"}, {"Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe3@m.co"}, {"Hanzo", "Hanzo0@m.co", "Hanzo1@m.co", "Hanzo3@m.co"}, {"Kevin", "Kevin0@m.co", "Kevin3@m.co", "Kevin5@m.co"}, {"Fern", "Fern0@m.co", "Fern1@m.co", "Fern5@m.co"}},
		},
	}

	for idx, tc := range cases {
		out := accountsMerge(tc.params)
		if !reflect.DeepEqual(out, tc.exp) {
			t.Errorf("case %d error: \n-- got: %v\n-- expected: %v", idx, out, tc.exp)
		}
	}
}
