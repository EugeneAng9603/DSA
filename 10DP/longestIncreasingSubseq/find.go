// O(nlogn), O(n)
// expected := []int{-24, 2, 3, 5, 6, 35}
// input := []int{5, 7, -24, 12, 10, 2, 3, 12, 5, 6, 35}
// this is optimised solution, for non-optimised, the usual DP way, see algoexpert
package main

import (
	"fmt"
)

func LongestIncreasingSubsequence(array []int) []int {
	// Write your code here.
	length := 0
	// len of this array show length of curr longest
	// stores index of latest number less than currValue
	// that ends an increasing sub-seq to itself
	// indices[i] means length of increasing sub-seq
	// that ends at indices[i]
	// e.g. [None, 2, 1, 3]
	// indices[2] is 1, means longest sub-seq that ends at array[1] = 7
	// is of length 2.... so on and so forth
	indices := make([]int, len(array)+1)
	seq := make([]int, len(array))
	for i := range array {
		seq[i] = -1
		indices[i] = -1
	}

	for i, num := range array {
		// binary search the value smaller than itself in indices[]
		// e.g. if get indices[2] = 5
		newLength := binarySearch(1, length, indices, array, num)
		fmt.Print(newLength, indices)
		fmt.Print("\n")
		// seq[i] always = indices[newLength-1], last smaller value in subseq
		seq[i] = indices[newLength-1]
		// update indices after binary searching the location
		indices[newLength] = i
		// update max length everytime
		length = max(length, newLength)
	}

	return buildSeq(array, seq, indices[length])
}

// return location of indices to be modified for each num in arr
// after BST using num and array[indices[mid]]
func binarySearch(start, end int, indices, array []int, num int) int {
	if start > end {
		return start
	}
	mid := (start + end) / 2
	if array[indices[mid]] > num {
		end = mid - 1
	} else {
		start = mid + 1
	}
	return binarySearch(start, end, indices, array, num)
}

func buildSeq(array, seq []int, index int) []int {
	result := []int{}
	for index != -1 {
		result = append(result, array[index])
		index = seq[index]
	}
	reverse(result)
	return result
}

// func max (a int, b ...int) int {
//     for _, num := range b {
//         if num > a {
//             a = num
//         }
//     }
//     return a
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverse(numbers []int) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}
