// wh, wh
package main

func TransposeMatrix(matrix [][]int) [][]int {
	// Write your code here.
	result := make([][]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		// MUST HAVE THIS STEP
		// To create a slice of each row for result before filling value
		newRow := make([]int, len(matrix))
		for j := 0; j < len(matrix); j++ {
			newRow[j] = matrix[j][i]
		}
		result[i] = newRow
	}
	return result
}

// func TransposeMatrix(matrix [][]int) [][]int {
// 	// Write your code here.
//     result := make([][]int, len(matrix[0]))
//     fmt.Print(result)
//     for i := 0; i < len(matrix[0]); i++ {
//         // newRow := make([]int, len(matrix))
//         for j := 0; j < len(matrix); j++ {
//             //result[i][j] = matrix[j][i]
//         }
//     }
// 	return result
// }
