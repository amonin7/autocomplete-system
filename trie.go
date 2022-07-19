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

//func (tr *Trie) Search(word string)

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
	autocompleteOneNode(*curNode, &result)
	return result
}

func autocompleteOneNode(curNode TrieNode, result *[]string) {
	if curNode.isEndingChar {
		*result = append(*result, curNode.word)
	}
	for _, child := range curNode.childrenChars {
		autocompleteOneNode(*child, result)
	}
}
