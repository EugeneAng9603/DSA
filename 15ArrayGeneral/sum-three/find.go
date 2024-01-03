// n2, n
package main

import (
	"sort"
)

func ThreeNumberSum(array []int, target int) [][]int {
	// Write your code here.
	sort.Ints(array)

	result := make([][]int, 0)
	for idx := 0; idx < len(array)-1; idx++ {
		left := idx + 1
		right := len(array) - 1

		for left < right {
			currSum := array[idx] + array[left] + array[right]

			if target == currSum {
				result = append(result, []int{array[idx], array[left], array[right]})
				left += 1
				right -= 1
			} else if currSum < target {
				left += 1
			} else {
				right -= 1
			}
		}
	}
	return result
}

// expected := [][]int{{-8, 2, 6}, {-8, 3, 5}, {-6, 1, 5}}
// output := ThreeNumberSum([]int{12, 3, 1, 2, -6, 5, -8, 6}, 0)
