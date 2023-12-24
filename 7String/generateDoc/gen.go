// characters = "abcd ef"
// document = "fed bca"
// true, freq of all char must be greater than all char in doc
package main

// import (
//
//	"fmt"
//
// )
func GenerateDocument(characters string, document string) bool {
	// Write your code here.
	table := map[rune]int{}

	for _, char := range characters {
		table[char] += 1
	}

	for _, char := range document {
		if table[char] == 0 {
			return false
		}
		table[char] -= 1
	}

	return true
}
