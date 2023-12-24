package main

import (
	"fmt"
)

// for every node in 2 for loops
// first we find UNVISITED neighbors! then bfs
// while doing dfs, pass down a count of 1
// rmb if all is 1 for 3x3, return 9

func RiverSizes(matrix [][]int) []int {
	// Write your code here.
	result := []int{}
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if visited[i][j] {
				continue
			}
			traverse(matrix, i, j, visited, &result)
		}
	}
	return result
}

// for every node, traverse each path BFS and record all river
// explore by checking the value == 0 or 1
func traverse(matrix [][]int, row, col int, visited [][]bool, result *[]int) {
	currSize := 0
	nodesToExplore := [][]int{{row, col}}
	for len(nodesToExplore) > 0 {
		currNode := nodesToExplore[0]
		nodesToExplore = nodesToExplore[1:]
		i, j := currNode[0], currNode[1]

		// check for visited first
		if visited[i][j] {
			continue
		}
		visited[i][j] = true

		// check if 0 or 1
		if matrix[i][j] == 0 {
			continue
		}
		currSize += 1

		// now add to nodesToExplore , check surrounding 4 cells
		neighbors := getUnvisited(i, j, matrix, visited)
		for _, neighbor := range neighbors {
			nodesToExplore = append(nodesToExplore, neighbor)
		}
	}

	// lastly add to result
	if currSize > 0 {
		*result = append(*result, currSize)
	}

}

func getUnvisited(i, j int, matrix [][]int, visited [][]bool) [][]int {
	unvisited := [][]int{}

	if i > 0 && !visited[i-1][j] {
		unvisited = append(unvisited, []int{i - 1, j})
	}
	if j > 0 && !visited[i][j-1] {
		unvisited = append(unvisited, []int{i, j - 1})
	}
	if i < len(matrix)-1 && !visited[i+1][j] {
		unvisited = append(unvisited, []int{i + 1, j})
	}
	if j < len(matrix[0])-1 && !visited[i][j+1] {
		unvisited = append(unvisited, []int{i, j + 1})
	}
	return unvisited
}

func main() {
	input := [][]int{
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
		{1, 0, 1, 0, 1},
	}

	fmt.Print(RiverSizes(input))
}
