// O(n+m), O(1)

// start from top right, if curr too big, row++
// if curr too small, col--
// cuz largest in first row is top right, if larger must be next row
// must check col first!!!
//       <--
//       |
//    <--
// so we check the range in the first row first
// e.g. 44 in <1000, then find correct row

// var matrix = [][]int{
// 	{1, 4, 7, 12, 15, 1000},
// 	{2, 5, 19, 31, 32, 1001},
// 	{3, 8, 24, 33, 35, 1002},
// 	{40, 41, 42, 44, 45, 1003},
// 	{99, 100, 103, 106, 128, 1004},
// }
// expected := []int{3, 3}
// output := SearchInSortedMatrix(matrix, 44)

package main

func SearchInSortedMatrix(matrix [][]int, target int) []int {
	// Write your code here.
	row := 0
	col := len(matrix[0]) - 1
	for row < len(matrix) && col >= 0 {
		if target < matrix[row][col] {
			col--
		} else if target > matrix[row][col] {
			row++
		} else {
			return []int{row, col}
		}
	}

	return []int{-1, -1}
}
