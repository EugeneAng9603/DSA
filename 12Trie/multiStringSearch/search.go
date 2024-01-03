// Can be better
// O(b^2 + ns) | O(b^2 + n) cuz we need b2 to put all word in trie
// put all word in big string in trie
// for every str in small, check
package main

import (
	"fmt"
)

func MultiStringSearch(bigString string, smallStrings []string) []bool {
	// Write your code here.
	// trie: this_is_a_big_string, his_is_a_big_string, isabigstring...
	trie := NewTrie(bigString)

	results := []bool{}
	for _, str := range smallStrings {
		// str is rune
		results = append(results, trie.Contains(str))
		fmt.Print(results)
	}

	return results
}

type SuffixTrie map[byte]SuffixTrie

// Feel free to add to this function.
func NewTrie(str string) SuffixTrie {
	trie := SuffixTrie{}
	for i := range str {
		trie.Add(str, i)
	}
	return trie
}

func (trie SuffixTrie) Add(str string, startIdx int) {
	// Write your code here.
	node := trie
	for j := startIdx; j < len(str); j++ {
		letter := str[j]
		if _, found := node[letter]; !found {
			node[letter] = SuffixTrie{}
		}
		node = node[letter]
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
	return true
}

// better optomised----------------------------------------
// O(ns+Bs) | O(ns)
// but small strings in trie: this*, yo*, is*, a*, bigger* etc...
// for every substring in big, find small string, if found put it in
// contained string (like this, is, a, string) set to true
// see algoexpert
