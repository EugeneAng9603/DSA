// O(n*n!), O(n*n!)
// Solution 1: BACKTRACK & ADD
package main

import (
	"fmt"
	"slices"
)

func GetPermutations(array []int) [][]int {
	// Write your code here.
	result := make([][]int, 0)
	backtrack(&result, []int{}, array)
	return result
}

func backtrack(result *[][]int, perm []int, array []int) {
	if len(perm) == len(array) {
		newPerm := make([]int, len(perm))
		copy(newPerm, perm)
		*result = append(*result, newPerm)
		return
	}
	for i := range array {
		if slices.Contains(perm, array[i]) {
			continue
		}
		perm = append(perm, array[i])
		backtrack(result, perm, array)
		perm = perm[:len(perm)-1]
	}
}

func GetPermutationsNoDup(array []int) [][]int {
	// Write your code here.
	result := make([][]int, 0)
	slices.Sort(array)
	used := make(map[int]bool)
	backtrackNoDup(&result, []int{}, array, used)
	return result
}

func backtrackNoDup(result *[][]int, perm []int, array []int, used map[int]bool) {
	if len(perm) == len(array) {
		newPerm := make([]int, len(perm))
		copy(newPerm, perm)
		*result = append(*result, newPerm)
		return
	}
	for i := range array {
		// if it's included or check duplicate
		// since we alr sorted, if curr == prev, and prev is not being used, means duplicate, skip this instance since already covered when i = prev
		// so we can guarantee we only go one time for 1,2,3,3 , when we reach the 4th 3, we wouldn't go back to use the 3rd 3 , so we skip it
		if used[i] || i > 0 && array[i] == array[i-1] && !used[i-1] {
			continue
		}
		perm = append(perm, array[i])
		used[i] = true
		backtrackNoDup(result, perm, array, used)
		perm = perm[:len(perm)-1]
		used[i] = false
	}
}

func main() {
	fmt.Print(len(GetPermutations([]int{1, 2, 3, 4}))) // 4!=24
	fmt.Print("\n")
	fmt.Print(len(GetPermutationsNoDup([]int{1, 2, 3, 3}))) // 4!/2!1!1! = 12
	fmt.Print("\n")
	fmt.Print(GetPermutationsNoDup([]int{1, 2, 3, 3}))
}

// --------------------------------------
// O(n*n!), O(n*n!)
// Solution 2: BACKTRACK & Swap
// package main

// import (
// 	"fmt"
// )

// func GetPermutations(array []int) [][]int {
// 	// Write your code here.
// 	result := make([][]int, 0)
// 	helper(0, array, &result)
// 	return result
// }

// func helper(level int, array []int, result *[][]int) {
// 	if level == len(array)-1 {
// 		// RMB to use newPerm to copy array's result
// 		// else it will only append [1,2,3] for all
// 		newPerm := make([]int, len(array))
// 		copy(newPerm, array)
// 		fmt.Printf("newPerm is : %v\n", array)
// 		*result = append(*result, newPerm)
// 		return
// 	}
// 	for j := level; j < len(array); j++ {
// 		swap(array, level, j)
// 		helper(level+1, array, result)
// 		swap(array, level, j)
// 	}
// }

// func swap(arr []int, i, j int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }

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
