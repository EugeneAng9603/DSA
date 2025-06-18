// mainly used to find patterns in a string.

// -- Write a function that takes in two strings and checks if the rst string
// contains the second one using the Knuth—Morris—Pratt algorithm. The function should return a boolean.

// Sample input:"aefoaefcdaefcdaed","aefcdaed" Sample output: True

// or question: https://leetcode.com/problems/longest-happy-prefix/solutions/547446/c-java-with-picture-incremental-hash-and-kmp/

// or question: find index of text := "ababcabcabababd", pattern := "ababd". Answer is 10.

// -- KMP Algorithm Explanation
// The Knuth-Morris-Pratt (KMP) algorithm is an efficient string matching algorithm.
// It improves upon the brute-force approach by avoiding unnecessary re-checking of characters in the text after a mismatch.

// -- Why KMP is Better Than Brute Force
// Feature	Brute Force	KMP
// Time Complexity	O(m * n)	O(m + n)
// Redundant Comparisons	Yes	No
// Preprocessing	No	Yes (Prefix table)
// Use case	Simple but slow	Fast and efficient

// Brute force tries all possible positions in the text and compares the pattern from scratch every time.

// KMP preprocesses the pattern using a prefix table (LPS - Longest Prefix Suffix), allowing it to "remember" where to resume after a mismatch.

// -- How KMP Works
// Build LPS (Longest Prefix Suffix) array: It tells us the longest proper prefix which is also a suffix for each sub-pattern.

// Search Phase: While matching, if a mismatch occurs, the LPS array tells us how many characters we can skip.

package main

import (
	"fmt"
)

// BuildLPS constructs the Longest Prefix Suffix array
func BuildLPS(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0
	i := 1

	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	fmt.Printf("LPS for pattern '%s': %v\n", pattern, lps)
	return lps
}

// KMPSearch searches for pattern in text and returns the indices of matches
func KMPSearch(text, pattern string) []int {
	lps := BuildLPS(pattern)
	var result []int

	i, j := 0, 0
	for i < len(text) {
		if text[i] == pattern[j] {
			i++
			j++
		}
		if j == len(pattern) {
			result = append(result, i-j)
			j = lps[j-1] // continue to search for next match
		} else if i < len(text) && text[i] != pattern[j] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return result
}

func main() {
	// text := "ababcabcabababd"
	// pattern := "ababd"
	text := "aaaaaaaaaaaaaaaaab"
	pattern := "aaaaaaaaaa"
	indices := KMPSearch(text, pattern)
	fmt.Println("Pattern found at indices:", indices)

	// Example usage of KMP_IfMatch
	// str := "aefoaefcdaefcdaed"
	// substr := "aefcdaed"
	str := "aaaaaaaaaaaaaaaaab"
	substr := "aaaaaaaaaa"
	if KMP_IfMatch(str, substr) {
		fmt.Printf("Substring '%s' found in string '%s'\n", substr, str)
	} else {
		fmt.Printf("Substring '%s' not found in string '%s'\n", substr, str)
	}
}

// algo expert
// knuthMorrisPrattAlgorithm returns true if substring is found in string
func KMP_IfMatch(str, substr string) bool {
	pattern := buildPattern(substr)
	return doesMatch(str, substr, pattern)
}

// buildPattern creates the prefix pattern array (equivalent to LPS array)
func buildPattern(substr string) []int {
	pattern := make([]int, len(substr))
	for i := range pattern {
		pattern[i] = -1
	}

	j := 0
	i := 1
	for i < len(substr) {
		if substr[i] == substr[j] {
			pattern[i] = j
			i++
			j++
		} else if j > 0 {
			j = pattern[j-1] + 1
		} else {
			i++
		}
	}
	fmt.Printf("LPS for pattern '%s': %v\n", substr, pattern)
	return pattern
}

// doesMatch uses the pattern to search for substring in string
func doesMatch(str, substr string, pattern []int) bool {
	i, j := 0, 0
	for i+len(substr)-j <= len(str) {
		if str[i] == substr[j] {
			if j == len(substr)-1 {
				return true
			}
			i++
			j++
		} else if j > 0 {
			j = pattern[j-1] + 1
		} else {
			i++
		}
	}
	return false
}
