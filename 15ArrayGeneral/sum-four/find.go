// n2, n2
// worst: n3, n2

// expected := [][]int{{7, 6, 4, -1}, {7, 6, 1, 2}}
// output := FourNumberSum([]int{7, 6, 4, -1, 1, 2}, 16)
package main

// 1. For every index [1 to n-1], break into left and right
// 2. Check for all left, add to table if not found, else append
// 3. Check for all right, if diff in table append, else skip, dont add new pairs
import (
	"sort"
)

// table:{13:[7,6], ...}
func FourNumberSum(array []int, target int) [][]int {
	// Write your code here.
	sort.Ints(array)
	table := map[int][][]int{}
	quads := make([][]int, 0)

	// use len(array) - 1 because last step is when j = last index, i = j-1
	for i := 1; i < len(array)-1; i++ {
		// to the right
		for j := i + 1; j < len(array); j++ {
			diff := target - array[i] - array[j]
			// add all in table that found
			if pairs, found := table[diff]; found {
				// if got 13: [7,6], [8,5], [9,4]
				// then keep adding
				for _, pair := range pairs {
					temp := append(pair, array[i], array[j])
					quads = append(quads, temp)
				}
				// else skip, dont add new pairs
			}
		}

		// to the left
		for k := 0; k < i; k++ {
			currSum := array[i] + array[k]
			table[currSum] = append(table[currSum], []int{array[k], array[i]})
		}

	}
	return quads
}
