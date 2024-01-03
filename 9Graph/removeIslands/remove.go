// O(Wh), O(Wh)
// remove all islands which are not on borders

// check border, if touch then not island
// for every border, dfs and set it to 2
// for all 1s, change non-islands to 2 in the same table,
// so no need create new matrix
// now, only islands are 1, the rest are 0, 2
// lastly, change all 1 to 0 to remove island,
// then 2 to 1 (all the non-islands 1's)
package main

func RemoveIslands(matrix [][]int) [][]int {
	// Write your code here.
	// visited := make([][]bool, len(matrix))
	// for i := range visited {
	//     // default false
	//     visited[i] = make([]bool, len(matrix[i]))
	// }

	// change non-islands to 2 first!!
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			// first, check border and connected to border
			// then change to 2
			rowIsBorder := i == 0 || i == len(matrix)-1
			colIsBorder := j == 0 || j == len(matrix[i])-1
			isBorder := rowIsBorder || colIsBorder

			if !isBorder {
				continue
			}

			if matrix[i][j] != 1 {
				continue
			}
			// if border and 1
			setTwo(matrix, i, j)
		}
	}

	// change all 2 to 1, 1 to 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			value := matrix[i][j]
			if value == 1 {
				matrix[i][j] = 0
			} else if value == 2 {
				matrix[i][j] = 1
			}
		}
	}
	return matrix
}

// no need visited, we traverse from borders 1 only
// check visited if needed etc, then set 2 first, then find neighbors
func setTwo(matrix [][]int, row, col int) {
	stack := [][]int{{row, col}}

	var curr []int
	for len(stack) > 0 {
		curr, stack = stack[len(stack)-1], stack[:len(stack)-1]
		currRow := curr[0]
		currCol := curr[1]

		matrix[currRow][currCol] = 2

		neighbors := getNeighbors(matrix, currRow, currCol)
		for _, neighbor := range neighbors {
			i, j := neighbor[0], neighbor[1]

			if matrix[i][j] != 1 {
				continue
			}
			stack = append(stack, neighbor)
		}

	}
}

func getNeighbors(matrix [][]int, row, col int) [][]int {
	neighbors := make([][]int, 0)
	rowlen := len(matrix)
	collen := len(matrix[row])

	if row > 0 {
		neighbors = append(neighbors, []int{row - 1, col}) // UP
	}
	if col > 0 {
		neighbors = append(neighbors, []int{row, col - 1}) // LEFT
	}
	if row < rowlen-1 {
		neighbors = append(neighbors, []int{row + 1, col}) // DOWN
	}
	if col < collen-1 {
		neighbors = append(neighbors, []int{row, col + 1}) // RIGHT
	}
	return neighbors
}

// input := [][]int{
// 	{1, 0, 0, 0, 0, 0},
// 	{0, 1, 0, 1, 1, 1},
// 	{0, 0, 1, 0, 1, 0},
// 	{1, 1, 0, 0, 1, 0},
// 	{1, 0, 1, 1, 0, 0},
// 	{1, 0, 0, 0, 0, 1},
// }
// expected := [][]int{
// 	{1, 0, 0, 0, 0, 0},
// 	{0, 0, 0, 1, 1, 1},
// 	{0, 0, 0, 0, 1, 0},
// 	{1, 1, 0, 0, 1, 0},
// 	{1, 0, 0, 0, 0, 0},
// 	{1, 0, 0, 0, 0, 1},
// }
