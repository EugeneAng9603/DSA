// Find longest substring with at most 2 distinct
// same logic, move end to find potential, then calculate

package main

import (
	"fmt"
)

func lengthOfLongestSubstringTwoDistinct(s string) (int, int) {
	head := 0
	result := 0
	table := map[byte]int{}
	start, end := 0, 0
	counter := 0
	for end < len(s) {
		if _, ok := table[s[end]]; !ok {
			table[s[end]] = 0
		}
		table[s[end]] += 1

		if table[s[end]] == 1 { // new char
			counter += 1
		}

		end += 1
		// move the start
		for counter > 2 {
			table[s[start]] -= 1
			if table[s[start]] == 0 {
				counter -= 1
			}
			start += 1
		}
		if end-start > result {
			result = end - start
			head = start
		}
	}
	return result, head
}

func main() {
	//fmt.Print(findAnagrams("babbbcbaebabacd", "aabc"))
	s := "AABCCCDEFFFFF"
	length, head := lengthOfLongestSubstringTwoDistinct(s)
	fmt.Print(s[head : head+length])
}
