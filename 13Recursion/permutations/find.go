// O(n*n!), O(n*n!)
// Solution 2: BACKTRACK & Swap
package main

import (
	"fmt"
)

func GetPermutations(array []int) [][]int {
	// Write your code here.
	result := make([][]int, 0)
	helper(0, array, &result)
	return result
}

func helper(level int, array []int, result *[][]int) {
	if level == len(array)-1 {
		// RMB to use newPerm to copy array's result
		// else it will only append [1,2,3] for all
		newPerm := make([]int, len(array))
		copy(newPerm, array)
		fmt.Printf("newPerm is : %v\n", array)
		*result = append(*result, newPerm)
		return
	}
	for j := level; j < len(array); j++ {
		swap(array, level, j)
		helper(level+1, array, result)
		swap(array, level, j)
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// ---------DFS---------------------------------------------------------
// Average: O(n*n!) , Upper: O(n2 * n!)
// O(n*n!)
// DFS, every level slice away 1, until bottom, at the result
// so keep removing 1 in another level , use a newPerm to add it back
// 1 [2 3 4] [1]
// 2 [3 4] [1 2]
// 3 [4] [1 2 3]
// 4 [] [1 2 3 4]
// 3 [3] [1 2 4]
// 4 [] [1 2 4 3] etc..
// package main

// import (
//     "fmt"
// )

// func GetPermutations(array []int) [][]int {
// 	// Write your code here.
//     result := [][]int{}
// 	helper(array, []int{}, &result, 1)
//     return result
// }

// func helper(arr, currPerm []int, perms *[][]int, level int) {
//     if len(arr) == 0 && len(currPerm) != 0 {
//         *perms = append(*perms, currPerm)
//         return
//     }
//     for i := range arr {
//         // newArr is current array of each level
//         // every level slice 1
//         // so [2,3] -> then [2] and [3],
//         // then [1,3] -> [1] and [3]
//         // slice itself, remaining put into newArr
//         newArr := make([]int, i)
//         copy(newArr, arr[:i])
//         // must separate arr into one, then append to newArr
//         newArr = append(newArr, arr[i+1:]...)

//         // newPerm is the passed down tracking result + curr_num
//         // when newPerm length = full, append this result
//         newPerm := make([]int, len(currPerm))
//         copy(newPerm, currPerm)
//         newPerm = append(newPerm, arr[i])
//         fmt.Print(level, newArr, newPerm)
//         fmt.Print("\n")
//         helper(newArr, newPerm, perms, level + 1)
//     }
// }
