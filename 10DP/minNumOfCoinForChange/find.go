// O(nd), O(n)
// n=7, denom=[1,5,10]
// output = 3 // 2x1 + 1x5

package main

import (
	"math"
)

func MinNumberOfCoinsForChange(n int, denoms []int) int {
	// Write your code here.

	table := make([]int, n+1)
	// make all value to max
	for i := range table {
		table[i] = math.MaxInt32
	}
	table[0] = 0

	for _, denom := range denoms {
		for amount := 1; amount < n+1; amount++ {
			if denom > amount {
				continue
			}
			table[amount] = min(table[amount-denom]+1, table[amount])
		}
	}

	if table[n] != math.MaxInt32 {
		return table[n]
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
