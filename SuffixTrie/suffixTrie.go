package main

import (
	"fmt"
)

type SuffixTrie map[byte]SuffixTrie

// Feel free to add to this function.
func NewSuffixTrie() SuffixTrie {
	trie := SuffixTrie{}
	return trie
}

func (trie SuffixTrie) PopulateSuffixTrieFrom(str string) {
	// Write your code here.
	for i := range str {
		node := trie
		for j := i; j < len(str); j++ {
			letter := str[j]
			if _, found := node[letter]; !found {
				node[letter] = NewSuffixTrie()
			}
			node = node[letter]
		}
		node['*'] = nil
	}
}

func (trie SuffixTrie) Contains(str string) bool {
	// Write your code here.
	node := trie
	for i := 0; i < len(str); i++ {
		letter := str[i]
		if _, found := node[letter]; !found {
			return false
		}
		node = node[letter]
	}
	_, found := node['*']
	return found
}

func main() {
	trie := NewSuffixTrie()
	string1 := "babc"
	trie.PopulateSuffixTrieFrom(string1)
	string2 := "ccc"
	fmt.Print(trie.Contains(string2))
}
