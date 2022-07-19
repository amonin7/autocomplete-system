package main

import "fmt"

type AutocompleteSystem struct {
	trie Trie
}

func SystemConstructor() *AutocompleteSystem {
	return &AutocompleteSystem{
		trie: *TrieConstructor(),
	}
}

func (acs *AutocompleteSystem) AddWords(words ...string) {
	for _, word := range words {
		acs.trie.Insert(word)
	}
}

func (acs *AutocompleteSystem) Autocomplete(query string, limit int) int {
	suggestions := acs.trie.Autocomplete(query, limit)
	if len(suggestions) == 0 {
		println(-1)
		return -1
	} else {
		fmt.Printf("%v", suggestions)
		return 0
	}
}
