// https://leetcode.com/problems/add-and-search-word-data-structure-design/
package add_and_search_word_data_structure_design

type WordDictionary struct {
	root *Node
}

type Node struct {
	next [27]*Node
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		root: &Node{},
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	n := this.root
	for i := 0; i < len(word); i++ {
		ch := word[i] - 'a'
		if n.next[ch] == nil {
			n.next[ch] = new(Node)
		}
		n = n.next[ch]
	}
	ch := 26 // EOW - end of word
	if n.next[ch] == nil {
		n.next[ch] = new(Node)
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.search([]byte(word), this.root)
}

func (this *WordDictionary) search(word []byte, n *Node) bool {
	if len(word) == 0 && n != nil && n.next[26] != nil {
		return true
	}
	if n == nil {
		return false
	}
	for i := 0; i < len(word); i++ {
		if word[i] == '.' {
			res := false
			for ch := 0; ch < 26; ch++ {
				res = res || this.search(word[1:], n.next[ch])
				if res {
					break
				}
			}
			return res
		} else {
			ch := word[i] - 'a'
			return this.search(word[1:], n.next[ch])
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
