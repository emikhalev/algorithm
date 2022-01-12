package add_and_search_word_data_structure_design

import (
	"testing"
)

func TestWordDictionary(t *testing.T) {
	wd := Constructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	assert(t, wd.Search("pad") == false)
	assert(t, wd.Search("bad"))
	assert(t, wd.Search(".ad"))
	assert(t, wd.Search("b.."))
}

func assert(t *testing.T, res bool) {
	if !res {
		t.Errorf("assert fails")
	}
}
