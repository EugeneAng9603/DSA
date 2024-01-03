// O(wnlogn), O(wn), w is number of words, n is len of longest word

// ["cinema", "iceman", "yo", "cat", "oy", "act", "tac"]
// output = [["cinema", "iceman"], ["act", "cat", "tac"], ["yo", "oy"]]
// we can sort each alphabetically before checking if anagrams are found
// map is string : arr of string as values
// then output all group of anagrams.

package main

import (
	"fmt"
	"sort"
)

func GroupAnagrams(words []string) [][]string {
	// Write your code here.
	anagrams := map[string][]string{}

	for _, word := range words {
		sortedWord := sortWord(word)
		anagrams[sortedWord] = append(anagrams[sortedWord], word)
	}
	//anagrams  {
	// oy: [ 'yo', 'oy' ],
	// act: [ 'act', 'tac', 'cat' ],
	// flop: [ 'flop', 'olfp' ],
	// foo: [ 'foo' ]
	// }
	result := [][]string{}
	for _, group := range anagrams {
		result = append(result, group)
	}
	return result
}

func sortWord(word string) string {
	wordBytes := []byte(word)

	sort.Slice(wordBytes, func(i, j int) bool {
		return wordBytes[i] < wordBytes[j]
	})
	fmt.Print(wordBytes)
	return string(wordBytes)
}
