// a bit different from template solution
// 395. Longest Substring with At Least K Repeating Characters
// Input: s = "aaabb", k = 3 , Output: 3
// Input: s = "ababbc", k = 2 Output: 5
// Given a string s and an integer k, return the length of the longest substring of s
// such that the frequency of each character in this substring is greater than or equal to k.
// https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/description/

package main

import (
	"fmt"
)

func longestSubstring(s string, k int) int {
	// find the number of different characters
	maxCharacterInSlidingWindow := getMaxDiffCharacter(s)
	// fmt.Println("Number of different character ", maxCharacterInSlidingWindow)
	result := 0
	for diffCharsInWindow := 1; diffCharsInWindow <= maxCharacterInSlidingWindow; diffCharsInWindow++ {
		/**
		each satisfying window will contain exactly diffCharsInWindwow
		character and all characters in window will be repeated at least k time
		 **/
		table := map[string]int{}
		currentDiff := 0 // to track how many distinct char we have until now
		left := 0
		right := 0
		satisfyingCharacters := 0 // track those char with freq = k

		for right < len(s) {
			// current window ranges from index left to right -1 (count right-1 as well )
			if currentDiff <= diffCharsInWindow {
				// we can process next character
				if table[s[right:right+1]] == 0 {
					currentDiff++
				}
				table[s[right:right+1]]++ // add one occurence
				if table[s[right:right+1]] == k {
					satisfyingCharacters++
				}
				right++
			} else {
				// we need to shrink the most left character in current window
				if table[s[left:left+1]] == k {
					satisfyingCharacters--
				}
				table[s[left:left+1]]--
				if table[s[left:left+1]] == 0 {
					currentDiff--
				}
				left++
			}
			if currentDiff == diffCharsInWindow && satisfyingCharacters == diffCharsInWindow {
				result = max(result, right-left)
			}
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxDiffCharacter(s string) int {
	table := map[string]bool{}
	for i := 0; i < len(s); i++ {
		table[s[i:i+1]] = true
	}
	return len(table)
}

// --------------------------------------------------------------------------------------------------
// solution not complete
// func longestSubstring(s string, k int) int {

// 	// head := 0
// 	result := 0
// 	table := map[byte]int{}
// 	start, end := 0, 0
// 	counter := 0
// 	for end < len(s) {
// 		if _, ok := table[s[end]]; !ok {
// 			table[s[end]] = 0
// 		}
// 		table[s[end]] += 1

// 		if table[s[end]] >= k { // new char
// 			counter += 1
// 		}

//			end += 1
//			// move the start
//			for counter != 0 {
//				table[s[start]] -= 1
//				if table[s[start]] < k {
//					counter -= 1
//				}
//				start += 1
//			}
//			if end-start > result {
//				result = end - start
//				// head = start
//			}
//		}
//		return result
//	}
func main() {

	fmt.Print(longestSubstring("aaabb", 3)) //3
	fmt.Print("\n")
	fmt.Print(longestSubstring("ababbc", 2)) //5

}
