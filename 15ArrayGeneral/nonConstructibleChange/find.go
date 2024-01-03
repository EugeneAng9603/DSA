package main

// input := []int{5, 7, 1, 1, 2, 3, 22}
// 	expected := 20
// cannot construct 20 of these coins

import (
	"sort"
)

func NonConstructibleChange(coins []int) int {
	// Write your code here.
	sort.Ints(coins)

	currChangeCreated := 0
	for _, coin := range coins {
		if coin > currChangeCreated+1 {
			return currChangeCreated + 1
		}
		currChangeCreated += coin
	}
	return currChangeCreated + 1
}
