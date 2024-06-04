// O(nd), O(n)
// n=7, denom=[1,5,10]
// output = 3 // 2x1 + 1x5
//           0 1 2 3 4 5 6 7
// table is [0 1 2 3 4 1 2 3]

package main

import (
	"fmt"
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

	fmt.Print(table)
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

func main() {
	fmt.Print(MinNumberOfCoinsForChange(7, []int{1, 5, 10}))
}
