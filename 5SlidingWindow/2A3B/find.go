// Find longest substring with exactly 2 A's and 3 B's

package main

import (
	"fmt"
)

// better way
func longestSubstringWith2A3B(s string) string {
	left := 0
	countA, countB := 0, 0
	maxLen := 0
	start := 0

	for right := 0; right < len(s); right++ {
		switch s[right] {
		case 'A':
			countA++
		case 'B':
			countB++
		}

		// Shrink window if we have too many A's or B's
		for countA > 2 || countB > 3 {
			switch s[left] {
			case 'A':
				countA--
			case 'B':
				countB--
			}
			left++
		}

		// Check for exactly 2 A's and 3 B's
		if countA == 2 && countB == 3 {
			if right-left+1 > maxLen {
				maxLen = right - left + 1
				start = left
			}
		}
	}

	if maxLen == 0 {
		return ""
	}
	return s[start : start+maxLen]
}

// func twoAthreeB(s string) (int, int) {
// 	head := 0
// 	result := 0
// 	table := map[byte]int{}
// 	start, end := 0, 0
// 	//counter := 0
// 	for end < len(s) {
// 		if _, ok := table[s[end]]; !ok {
// 			table[s[end]] = 0
// 		}
// 		table[s[end]] += 1

// 		// lesser than 2A3B, more than 2A or more than 3B. Then update max. and Move left to correct position using
// 		if table['A'] == 2 && table['B'] == 3 {
// 			head = start
// 			result = max(result, end-start+1)
// 			end += 1
// 			continue
// 		}

// 		if table['B'] > 3 {
// 			for s[start] != 'B' {
// 				table[s[start]] -= 1
// 				start += 1
// 			}
// 			table[s[start]] -= 1
// 			start += 1
// 			end += 1
// 			continue
// 		} else if table['A'] > 2 {
// 			for s[start] != 'A' {
// 				table[s[start]] -= 1
// 				start += 1
// 			}
// 			table[s[start]] -= 1
// 			start += 1
// 			end += 1
// 			continue
// 		}

// 		end += 1
// 		fmt.Print(s[head : head+result])
// 		fmt.Print("\n")
// 	}
// 	return result, head
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//fmt.Print(findAnagrams("babbbcbaebabacd", "aabc"))
	s := "AABBBCCBCDEEEEEBABFFFFFAB" //19
	// s := strings.ToUpper("babbbcbaebabacd") //7
	// length, head := twoAthreeB(s)
	// fmt.Print(s[head:head+length], head, length)

	fmt.Print(longestSubstringWith2A3B(s)) // should print "AABBBCCBCDEEEEEBABFFFFFAB" or similar with 2 A's and 3 B's
}
