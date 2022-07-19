package main

/*
	w -> o -> r -> d

	d -> {empty map, true, "word"}
	o -> {non-empty map, false, ""}
*/

type TrieNode struct {
	childrenChars map[rune]*TrieNode
	isEndingChar  bool
	word          string
}

func TrieNodeConstructor() *TrieNode {
	return &TrieNode{
		childrenChars: make(map[rune]*TrieNode),
		isEndingChar:  false,
		word:          "",
	}
}
