// O(mn), O(mn)
// str1 = "abc" , str2 = "yabd"
// output = 2

// Initialize table 012345, 012345 ...
// for 1 -> len + 1, do the checking
package main

import (
	"fmt"
)

func LevenshteinDistance(a, b string) int {
	// Write your code here.
	table := make([][]int, len(b)+1)
	// initialise table
	// add col to every row
	for y := range table {
		table[y] = make([]int, len(a)+1)
		for x := range table[y] {
			// add value to every cell
			table[y][x] = x
		}
	}
	// first col must be 0,1,2,3 ..., not all 0 following col value
	for i := 1; i < len(b)+1; i++ {
		table[i][0] = table[i-1][0] + 1
	}

	fmt.Print(table)
	// main logic
	for i := 1; i < len(b)+1; i++ {
		for j := 1; j < len(a)+1; j++ {
			// if same char
			if b[i-1] == a[j-1] {
				table[i][j] = table[i-1][j-1]
			} else {
				table[i][j] = 1 + min(table[i-1][j], table[i-1][j-1], table[i][j-1])
			}
		}
	}

	return table[len(b)][len(a)]
}

func min(args ...int) int {
	curr := args[0]
	for _, num := range args {
		if num < curr {
			curr = num
		}
	}
	return curr
}
