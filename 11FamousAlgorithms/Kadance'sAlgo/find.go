// expected := 19
// 	output := KadanesAlgorithm([]int{3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4})
// 	require.Equal(t, expected, output)

// O(n), O(1)

// for every value, check 2 things
// 1. currSum, choose currVal or currVal + currSum
// 2. update result
package main

// import (
//     "math"
// )

func KadanesAlgorithm(array []int) int {
	// Write your code here.
	result := array[0]
	currSum := array[0]
	for i := 1; i < len(array); i++ {
		currSum = max(array[i], currSum+array[i])
		result = max(result, currSum)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
