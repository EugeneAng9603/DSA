// O(b+s), O(b+s), b is len of larger string, s is smaller

// big = "bcd$ef$axb$c$"
// small = "$$abf"
// outout = "f$axb$"
package main

import (
	"fmt"
)

func SmallestSubstringContaining(bigString, smallString string) string {
	if len(smallString) > len(bigString) {
		return ""
	}

	table := map[byte]int{}
	for i := 0; i < len(smallString); i++ {
		table[smallString[i]] += 1
	}

	start, end, head := 0, 0, 0
	// head is the head of result substring
	counter := len(table)
	longest := 1000000 //math.MaxInt32
	for end < len(bigString) {
		// found any of "AABC", then update counter accordingly
		if _, ok := table[bigString[end]]; ok {
			table[bigString[end]] -= 1 //reduce frequency, can reach -1
			if table[bigString[end]] == 0 {
				counter -= 1 // means alredy have all "A", met req. so counter -1
			}
		}
		// can put at the end of code
		end += 1
		fmt.Print(start, end, table)
		fmt.Print("\n")
		// move the start
		for counter == 0 { //means currently have substring including all "AABC"
			// this part is to update start, first check if start in "AABC",
			// if yes, counter++ , no need update start for next

			if _, ok := table[bigString[start]]; ok {
				table[bigString[start]] += 1
				if table[bigString[start]] > 0 {
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

	return bigString[head : head+longest]

}

func main() {
	big := "bcd$ef$axb$c$"
	small := "$$abf"
	fmt.Print(SmallestSubstringContaining(big, small))
}
