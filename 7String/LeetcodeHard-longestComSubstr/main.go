// Given a string s, consider all duplicated substrings: (contiguous) substrings of s that occur 2 or more times. The occurrences may overlap.

// Return any duplicated substring that has the longest possible length. If s does not have a duplicated substring, the answer is "".

// Example 1:

// Input: s = "banana"
// Output: "ana"
// Example 2:

// Input: s = "abcd"
// Output: ""

// ✅ Pros and Cons of Rabin-Karp vs Suffix Array:
// Technique	Time Complexity	Pros	Cons
// Suffix Array	O(n log n)	Very precise, lex order useful	More complex to implement
// Rabin-Karp	O(n log n)	Easier, efficient for this use case	Risk of hash collisions (low)

package main

import "fmt"

func main() {
	tests := []struct {
		input    string
		expected string
	}{
		{"banana", "ana"},
		{"abcdef", ""},
		{"aabbaabbaabb", "aabbaabb"},
		{"abcabcabc", "abcabc"},
		{"mississippi", "issi"},
		{"", ""},
	}

	for _, test := range tests {
		result := FindLongestRepeatedSubstringRabin(test.input)
		result2 := FindLongestRepeatedSubstringSuffix(test.input)
		fmt.Printf("Input: %-15s | Output: %-10s | Expected: %-10s | Pass: %v\n",
			test.input, result, test.expected, result == test.expected)
		fmt.Printf("Input: %-15s | Output: %-10s | Expected: %-10s | Pass: %v\n",
			test.input, result2, test.expected, result2 == test.expected)
	}
}
