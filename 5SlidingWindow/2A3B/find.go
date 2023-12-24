// Find longest substring with exactly 2 A's and 3 B's

package main

import (
	"fmt"
)

func twoAthreeB(s string) (int, int) {
	head := 0
	result := 0
	table := map[byte]int{}
	start, end := 0, 0
	//counter := 0
	for end < len(s) {
		if _, ok := table[s[end]]; !ok {
			table[s[end]] = 0
		}
		table[s[end]] += 1

		// lesser than 2A3B, more than 2A or more than 3B. Then update max. and Move left to correct position using
		if table['A'] == 2 && table['B'] == 3 {
			head = start
			result = max(result, end-start+1)
			end += 1
			continue
		}

		if table['B'] > 3 {
			for s[start] != 'B' {
				table[s[start]] -= 1
				start += 1
			}
			table[s[start]] -= 1
			start += 1
			end += 1
			continue
		} else if table['A'] > 2 {
			for s[start] != 'A' {
				table[s[start]] -= 1
				start += 1
			}
			table[s[start]] -= 1
			start += 1
			end += 1
			continue
		}
		end += 1
		fmt.Print(s[head : head+result])
		fmt.Print("\n")
	}
	return result, head
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//fmt.Print(findAnagrams("babbbcbaebabacd", "aabc"))
	s := "AABBBCCBCDEEEEEBABFFFFFAB"
	//s := strings.ToUpper("babbbcbaebabacd")
	length, head := twoAthreeB(s)
	fmt.Print(s[head:head+length], head, length)
}
