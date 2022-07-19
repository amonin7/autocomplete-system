package main

type Trie struct {
	root *TrieNode
}

func TrieConstructor() *Trie {
	return &Trie{
		root: TrieNodeConstructor(),
	}
}

func (tr *Trie) Insert(word string) {
	curNode := tr.root
	for _, character := range word {
		childNode, ok := curNode.childrenChars[character]
		if !ok {
			childNode = TrieNodeConstructor()
			curNode.childrenChars[character] = childNode
		}
		curNode = childNode
	}
	curNode.isEndingChar = true
	curNode.word = word
}

func (tr *Trie) Search(word string) bool {
	curNode := tr.root
	for _, character := range word {
		childNode, ok := curNode.childrenChars[character]
		if !ok {
			return false
		}
		curNode = childNode
	}
	return curNode.isEndingChar
}

func (tr *Trie) StartsWith(prefix string) bool {
	curNode := tr.root
	for _, character := range prefix {
		childNode, ok := curNode.childrenChars[character]
		if !ok {
			return false
		}
		curNode = childNode
	}
	return true
}

func (tr *Trie) Autocomplete(query string, limit int) []string {
	curNode := tr.root
	for _, character := range query {
		childNode, ok := curNode.childrenChars[character]
		if !ok {
			// there is no such prefix in the trie
			return []string{}
		}
		curNode = childNode
	}

	var result []string
	autocompleteOneNode(*curNode, &result, limit)
	return result
}

func autocompleteOneNode(curNode TrieNode, result *[]string, limit int) {
	if limit != 0 && len(*result) == limit {
		return
	}
	if curNode.isEndingChar {
		*result = append(*result, curNode.word)
	}
	if limit != 0 && len(*result) == limit {
		return
	}
	for _, child := range curNode.childrenChars {
		autocompleteOneNode(*child, result, limit)
	}
}
