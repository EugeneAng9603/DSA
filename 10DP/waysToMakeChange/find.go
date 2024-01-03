// O(nd), O(n)

// ways[0] = 1, count up to ways[n+1]
// for every denom, count its ways up to n+1
// by using ways[amount] += ways[amount-denom]
// meaning currWay + ways of 6-5
// e.g. denom = 5, ways[6] += ways[1]
// n=6, denom = [1,5]
// output = 2 // 1x1+1x5 , 6x1

package main

func NumberOfWaysToMakeChange(n int, denoms []int) int {
	// Write your code here.
	ways := make([]int, n+1)
	ways[0] = 1
	for _, denom := range denoms {
		for amount := 1; amount < n+1; amount++ {
			if amount >= denom {
				ways[amount] += ways[amount-denom]
			}
		}
	}
	return ways[n]
}
