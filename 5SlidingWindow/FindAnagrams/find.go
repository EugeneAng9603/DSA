// Given two strings s and p, return an array of all the start indices of p's anagrams in s.
// You may return the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.

// Example 1:

// Input: s = "cbaebabacd", p = "abc"
// Output: [0,6]
// Explanation:
// The substring with start index = 0 is "cba", which is an anagram of "abc".
// The substring with start index = 6 is "bac", which is an anagram of "abc".
// Example 2:

// Input: s = "abab", p = "ab"
// Output: [0,1,2]
// Explanation:
// The substring with start index = 0 is "ab", which is an anagram of "ab".
// The substring with start index = 1 is "ba", which is an anagram of "ab".
// The substring with start index = 2 is "ab", which is an anagram of "ab".

// 1: we start by moving end to find potential substring with "aabc", if found,
// 2: move start to find , it will be found when end-start == length of p
// 3: if not found until s[start] > 0, counter ++ cuz curr char is skipped, means no result, then stop
// and repeat step 1

package main

import (
	"fmt"
)

func findAnagrams(s string, p string) []int {
	result := []int{}
	if len(p) > len(s) {
		return result
	}

	table := map[byte]int{}
	for i := 0; i < len(p); i++ {
		table[p[i]] += 1
	}

	start, end := 0, 0
	counter := len(table)
	//longest := 1000000
	for end < len(s) {
		// found any of "AABC", then update counter accordingly
		if _, ok := table[s[end]]; ok {
			table[s[end]] -= 1      //reduce frequency. can reach -1
			if table[s[end]] == 0 { // while moving end, whenever we found s[end] == 0, means fulfilled for 1 char
				counter -= 1 // so counter - 1
			}
		}
		// can put at the end of code

		fmt.Print(table, "counter:", counter, start, string(s[end]), end)
		fmt.Print("\n")
		end += 1
		// move the start
		for counter == 0 { //means currently have substring including all "AABC", fulfilled for all char
			// be it from -1 to 0 or 1 to 0
			// this part is to update start, first check if start in "AABC",
			// if yes, counter++ , no need update start for next
			if _, ok := table[s[start]]; ok {
				table[s[start]] += 1
				if table[s[start]] > 0 {
					// means this substring != len(p) although we have the "aabc" in it, counter ++, abort
					counter += 1
				}
			}
			// update min
			if end-start == len(p) {
				result = append(result, start)
			}

			fmt.Print("For start:", table, "counter:", counter, string(s[start]), start, end, result)
			fmt.Print("\n")
			start += 1
		}
	}
	return result
}

func main() {
	//fmt.Print(findAnagrams("babbbcbaebabacd", "aabc"))
	fmt.Print(findAnagrams("bbbbcbaaebabacd", "aabc"))
}
