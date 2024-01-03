// n,n
package main

func SpiralTraverse(array [][]int) []int {
	// Write your code here.
	result := make([]int, 0)
	helper(array, 0, len(array)-1, 0, len(array[0])-1, &result)
	return result
}

func helper(arr [][]int, startRow, endRow, startCol, endCol int, result *[]int) {
	// stop case
	if startRow > endRow || startCol > endCol {
		return
	}
	for j := startCol; j <= endCol; j++ {
		*result = append(*result, arr[startRow][j])
	}
	for i := startRow + 1; i <= endRow; i++ {
		*result = append(*result, arr[i][endCol])
	}
	for j := endCol - 1; j >= startCol; j-- {
		// if uneven matrix, handle duplicate
		if startRow == endRow {
			break
		}
		*result = append(*result, arr[endRow][j])
	}
	for i := endRow - 1; i > startCol; i-- {
		if startCol == endCol {
			break
		}
		*result = append(*result, arr[i][startCol])
	}

	helper(arr, startRow+1, endRow-1, startCol+1, endCol-1, result)
}

// matrix := [][]int{
// 	{1, 2, 3, 4},
// 	{12, 13, 14, 5},
// 	{11, 16, 15, 6},
// 	{10, 9, 8, 7},
// }
// expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
