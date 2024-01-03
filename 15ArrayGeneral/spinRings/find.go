// n, 1

// Recursive using spiral traverse logic
package main

// import (
//     "fmt"
// )

func SpinRings(array [][]int) {
	result := make([]int, 0)
	//final := make([][]int, 0)
	helper(array, 0, len(array)-1, 0, len(array[0])-1, &result)
	//return array
}

func helper(arr [][]int, startRow, endRow, startCol, endCol int, result *[]int) {
	// stop case
	if startRow > endRow || startCol > endCol {
		return
	}

	prev := arr[startRow][startCol]
	for j := startCol + 1; j <= endCol; j++ {
		temp := arr[startRow][j]
		arr[startRow][j] = prev
		prev = temp
	}
	for i := startRow + 1; i <= endRow; i++ {
		temp := arr[i][endCol]
		arr[i][endCol] = prev
		prev = temp
	}
	for j := endCol - 1; j >= startCol; j-- {
		// if uneven matrix, handle duplicate
		if startRow == endRow {
			break
		}
		temp := arr[endRow][j]
		arr[endRow][j] = prev
		prev = temp
	}
	for i := endRow - 1; i > startCol; i-- {
		if startCol == endCol {
			break
		}
		temp := arr[i][startCol]
		arr[i][startCol] = prev
		prev = temp
	}
	arr[startRow][startCol] = prev
	helper(arr, startRow+1, endRow-1, startCol+1, endCol-1, result)
}

// array := [][]int{
// 	{1, 2, 3, 4},
// 	{5, 6, 7, 8},
// 	{9, 10, 11, 12},
// 	{13, 14, 15, 16},
// }
// expected := [][]int{
// 	{5, 1, 2, 3},
// 	{9, 10, 6, 4},
// 	{13, 11, 7, 8},
// 	{14, 15, 16, 12},
// }
// SpinRings(array)
