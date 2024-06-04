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
	// that ends at array[i]
	// indices[1] means last index that forms length of 1
	// indices[5] means last index that forms length of 5
	// e.g. [None, 2, 1, 3,...]
	//  arr= [5,   7,-24,12,...]
	// indices[2] is 1, means longest sub-seq that ends at array[1] = 7 is of length 2.... so on and so forth

	// for every num in arr, we want to find a smaller value with largest length
	indices := make([]int, len(array)+1)
	seq := make([]int, len(array))
	for i := range array {
		seq[i] = -1
		indices[i] = -1
	}

	for i, num := range array {
		// newLength is the max length we can have until num, so update indices
		newLength := binarySearch(1, length, indices, array, num)
		// seq[i] always = indices[newLength-1], last smaller value in subseq
		seq[i] = indices[newLength-1]
		// update indices after binary searching the location
		indices[newLength] = i
		fmt.Print("new: ", indices, "currNum is:", num, " length is: ", length, " newLength is: ", newLength)
		fmt.Print("\n")
		// fmt.Print("seq: ", seq)
		fmt.Print("\n")
		// update max length everytime
		length = max(length, newLength)
	}

	return buildSeq(array, seq, indices[length])
}

// search for a num that yields a longer length
// if mid < num, means we found a good candidate, we can try to find longer length, so check right half
// if cant find another smaller value, then that's the maxLength until num
// if mid > num, means we can only find shorter length, so go left
// cuz if mid cant, right of mid (longer length) also cant

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

func main() {
	input := []int{5, 7, -24, 12, 10, 2, 3, 12, 5, 6, 35, -1, -2, 36}
	fmt.Print("input", input, "\n")
	fmt.Print(LongestIncreasingSubsequence(input))
}

// --------------------------------------------------------------------------------------------
// func lengthOfLIS(nums []int) int {
//     dp := make([]int, len(nums))

//     res := 0
//     for i := 0; i < len(nums); i++ {
//         dp[i] = 1
//         for j := 0; j < i; j++ {
//             if nums[j] < nums[i] {
//                 dp[i] = max(dp[i], dp[j] + 1)
//             }
//         }
//         res = max(res, dp[i])
//     }

//     return res
// }

// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }
