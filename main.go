package main

/*
We are given a Trie with a set of strings stored in it.
Now the user types in a prefix of his search query,
we need to give him all recommendations to auto-complete his query based on the strings stored in the Trie.
We assume that the Trie stores past searches by the users.

For example if the Trie store {“abc”, “abcd”, “aa”, “abbbaba”} and
the User types in “ab” then he must be shown {“abc”, “abcd”, “abbbaba”}.

1. Search for the given query using the standard Trie search algorithm.

2. If the query prefix itself is not present, return -1 to indicate the same.

3. If the query is present and is the end of a word in Trie, print query.
This can quickly be checked by seeing if the last matching node has isEndWord flag set.
We use this flag in Trie to mark the end of word nodes for purpose of searching.

4. Else recursively print all nodes under a subtree of last matching node.
*/

func main() {
	acs := SystemConstructor()
	acs.AddWords("hello", "dog", "hell", "cat", "a", "hel", "help", "helps", "helping")
	acs.Autocomplete("c", 0)
}
