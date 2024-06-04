// nm, nm

// add all to trie first
// use a count in the recursive isSpecial function, if more than 2 then true

// strings := []string{
// 	"foobarbaz",
// 	"foo",
// 	"bar",
// 	"foobarfoo",
// 	"baz",
// 	"foobaz",
// 	"foofoofoo",
// 	"foobazar",
// }
// actual := SpecialStrings(strings)
// expected := []string{"foobarbaz", "foobarfoo", "foobaz", "foofoofoo"}
// require.Equal(t, expected, actual)

package main

import (
	"fmt"
)

// trie:
//  &{*, {t:{h:{i:{s:{}}}, e:{s:{t:{}}}, ... }}

func SpecialStrings(strings []string) []string {
	// Write your code here.
	trie := newTrie()
	for _, str := range strings {
		trie.insert(str)
	}

	fmt.Print(trie)
	result := []string{}
	// for every element in strings, check isSpecial , if true append
	for _, str := range strings {
		if isSpecial(str, trie.root, 0, 0, trie) {
			result = append(result, str)
		}
	}
	return result
}

// Recursively check all characters in the element
// if already at end of element, check if valid (must be endSymbol)
// and count > 2
// if complete(means reach endSymbol of trie), e.g. foo
// keep checking for nextNode and idx + 1 with count + 1
// e.g. for foobaz, if foo valid, check count = 1
// if complete then go back to root to check baz
func isSpecial(str string, node trieNode, idx, count int, t *trie) bool {
	char := rune(str[idx])
	nextTrieNode, found := node[char]
	if !found {
		return false
	}

	atEndOfString := idx == len(str)-1
	_, isComplete := nextTrieNode[t.endSymbol]

	if atEndOfString {
		return isComplete && count+1 >= 2
	}
	if isComplete {
		restIsSpecial := isSpecial(str, t.root, idx+1, count+1, t)
		if restIsSpecial {
			return true
		}
	}

	return isSpecial(str, nextTrieNode, idx+1, count, t)
}

type trie struct {
	endSymbol rune
	root      trieNode
}

type trieNode map[rune]trieNode

func newTrie() *trie {
	return &trie{
		endSymbol: '*',
		root:      trieNode{},
	}
}

func (t *trie) insert(str string) {
	currNode := t.root
	for i := 0; i < len(str); i++ {
		currStr := str[i]
		child, found := currNode[rune(currStr)]
		if !found {
			child = trieNode{}
			currNode[rune(currStr)] = child
		}
		currNode = child
	}
	currNode[t.endSymbol] = trieNode{}
}

func main() {
	strings := []string{
		"foobarbaz",
		"foo",
		"bar",
		"foobarfoo",
		"baz",
		"foobaz",
		"foofoofoo",
		"foobazar",
	}

	fmt.Print(SpecialStrings(strings))
}
