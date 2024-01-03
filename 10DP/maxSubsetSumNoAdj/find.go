// O(n), O(1)
// res := MaxSubsetSumNoAdjacent([]int{75, 105, 120, 75, 90, 135})
// 	require.Equal(t, res, 330) // 75+120+135

package main

func MaxSubsetSumNoAdjacent(array []int) int {
	// Write your code here.

	if len(array) == 0 {
		return 0
	} else if len(array) == 1 {
		return array[0]
	}
	first := array[0]
	second := max(array[0], array[1])
	for i := 2; i < len(array); i++ {
		currSum := max(second, first+array[i])
		first, second = second, currSum
	}

	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
