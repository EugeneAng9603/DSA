// O(n * 2^n), O(n * 2^n)
package main

// output := Powerset([]int{1, 2, 3})
// 	require.Contains(t, output, []int{})
// 	require.Contains(t, output, []int{1})
// 	require.Contains(t, output, []int{2})
// 	require.Contains(t, output, []int{1, 2})
// 	require.Contains(t, output, []int{3})
// 	require.Contains(t, output, []int{1, 3})
// 	require.Contains(t, output, []int{2, 3})
// 	require.Contains(t, output, []int{1, 2, 3})
// 	require.Len(t, output, 8)
// }

import (
	"fmt"
)

func Powerset(array []int) [][]int {
	// Write your code here.
	powerset := [][]int{}
	helper(array, []int{}, &powerset, 0)
	return powerset
}

func helper(arr []int, currArr []int, powerset *[][]int, start int) {
	// MUST copy into temp to add to powerset
	temp := make([]int, len(currArr))
	copy(temp, currArr)
	*powerset = append(*powerset, temp)
	fmt.Print(*powerset, start)
	fmt.Print("\n")
	if start >= len(arr) {
		return
	}
	for i := start; i < len(arr); i++ {
		currArr = append(currArr, arr[i])
		helper(arr, currArr, powerset, i+1)
		currArr = currArr[:len(currArr)-1]
	}
}

// option2 backtrack ----------------------------------------
// BACKTRACK
// package main

// import (
//     "fmt"
// )

// func Powerset(array []int) [][]int {
// 	// Write your code here.
//     results := [][]int{}
//     helper(0, []int{}, array, &results)
// 	return results
// }

// func helper(idx int, path []int, array []int, results *[][]int) {

//     // this is a new subset, all the time (even empty array at beginning)
//     // so we should copy path, and add to result
//     pathTemp := make([]int, len(path))
//     copy(pathTemp, path)
//     *results = append(*results, pathTemp)
//     fmt.Print(idx, *results)
//     fmt.Print("\n")
//     // go from start to end (because we want no repeaters)
//     for i := idx; i < len(array); i++ {
//         // take this number, recurse that path,
//         // and then pop cause done searching
//         path = append(path, array[i])
//         helper(i+1, path, array, results)
//         // once done, reached reached [1,2,3,4]
//         // remove 4 and i move to 3 then keep doing
//         // but nth else to add for 1,2,3
//         // move to 1,2, add 4, then back to 1, idx at 3
//         path = path[:len(path)-1]
//     }
// }
