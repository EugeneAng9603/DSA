// Leetcode 53

package main

func maxSubArray(array []int) int {
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
