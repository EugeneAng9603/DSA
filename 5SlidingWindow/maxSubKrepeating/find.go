// a bit different from template solution
// 395. Longest Substring with At Least K Repeating Characters
// Input: s = "aaabb", k = 3 , Output: 3
// Input: s = "ababbc", k = 2 Output: 5
// Given a string s and an integer k, return the length of the longest substring of s
// such that the frequency of each character in this substring is greater than or equal to k.
// https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/description/

package main

import "fmt"

func find(s string, k int) int {

}

func main() {
	s := "ababbc"
	k := 2
	fmt.Print(find(s, k))
}
