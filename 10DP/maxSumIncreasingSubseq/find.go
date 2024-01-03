// O(n2), O(n)
// use an array to store the maxsum up to itself
// intput := []int{10, 70, 20, 30, 50, 11, 30}
// 	require.Equal(t, 110, outputSum)
// 	require.Equal(t, []int{10, 20, 30, 50}, outputSequence)

package main

import (
	"math"
)

func MaxSumIncreasingSubsequence(array []int) (int, []int) {
	// Write your code here.

	// store last smaller value's index
	seq := make([]int, len(array))

	// maxSum up to itslef
	sums := make([]int, len(array))

	// assign all value of seq to math.MinInt32 first
	// assume all maxSums is just itself without adding any other
	for i := range seq {
		seq[i] = math.MinInt32
		sums[i] = array[i]
	}

	// use maxSumIndex to build the array later
	// build the result starting from maxSumIndex
	maxSumIndex := 0
	for i := 0; i < len(array); i++ {
		for j := 0; j < i; j++ {
			otherNum := array[j] // all number before itself
			currNum := array[i]
			// arr[j] < arr[i] means can be increasing subseq
			// then check larger of sums[j]+arr[i] , sums[i]
			// if true, update sum and seq
			if otherNum < currNum && sums[j]+currNum >= sums[i] {
				sums[i] = sums[j] + currNum
				seq[i] = j
			}
		}
		// update maxSumIndex
		if sums[i] > sums[maxSumIndex] {
			maxSumIndex = i
		}
	}

	maxSum := sums[maxSumIndex]
	resultArr := buildSequence(array, seq, maxSumIndex)
	return maxSum, resultArr
}

func buildSequence(array, seq []int, index int) []int {
	//result := make([]int, 0)
	result := []int{}
	for index != math.MinInt32 {
		result = append(result, array[index])
		index = seq[index]
	}
	reverse(result)
	return result
}

// func reverse(arr []int) {
//     for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
//         arr[i], arr[j] = arr[j], arr[i]
//     }
// }

func reverse(arr []int) {
	start := 0
	end := len(arr) - 1
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
