// correct solution using template
// same as anagrams

package main

import (
	"fmt"
)

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	table := map[byte]int{}
	for i := 0; i < len(t); i++ {
		table[t[i]] += 1
	}

	start, end, head := 0, 0, 0
	// head is the head of result substring
	counter := len(table)
	longest := 1000000
	for end < len(s) {
		// found any of "AABC", then update counter accordingly
		if _, ok := table[s[end]]; ok {
			table[s[end]] -= 1 //reduce frequency, can reach -1
			if table[s[end]] == 0 {
				counter -= 1
			}
		}
		// can put at the end of code
		end += 1

		// move the start
		for counter == 0 { //means currently have substring including all "AABC"
			// this part is to update start, first check if start in "AABC",
			// if yes, counter++ , no need update start for next
			if _, ok := table[s[start]]; ok {
				table[s[start]] += 1
				if table[s[start]] > 0 {
					//counter++, so no need update "start" now, go to next end, new part of substring
					counter += 1
				}
			}
			// update min
			if end-start < longest {
				longest = end - start
				head = start
			}
			start += 1
		}

	}
	if longest == 1000000 {
		return ""
	}

	return s[head : head+longest]

}

func main() {
	fmt.Print(minWindow("ADOBECODEBANC", "ABC"))
}
