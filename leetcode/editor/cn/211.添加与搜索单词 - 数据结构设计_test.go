package cn

import "testing"

func TestWordDictionary_Search(t *testing.T) {

	trie := Constructor()
	trie.AddWord("bad")
	trie.AddWord("dad")
	trie.AddWord("mad")
	trie.AddWord("add")

	println(trie.Search("bad"))
	println(trie.Search("..d"))
}
