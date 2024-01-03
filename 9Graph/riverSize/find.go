// O(wh), O(wh)

// for every '1', add the unvisited neighbor of '1'
// then BFS for 1
// if 'H' shape, considered 7
// if 7x7, considered 49

package main

import (
	"fmt"
)

func RiverSizes(matrix [][]int) []int {
	// Write your code here.
	result := make([]int, 0)
	//result := []int{}
	visited := make([][]bool, len(matrix))
	for i := range visited {
		// default false
		visited[i] = make([]bool, len(matrix[i]))
	}

	for i := range matrix {
		for j := range matrix[i] {
			if visited[i][j] {
				continue
			}
			result = dfs(i, j, matrix, visited, result)
		}
	}
	return result
}

// do all checking, currSize++, lastly add neighbors
func dfs(row, col int, matrix [][]int, visited [][]bool, result []int) []int {
	currSize := 0
	nodesToExplore := [][]int{{row, col}}
	for len(nodesToExplore) > 0 {
		currNode := nodesToExplore[0]
		nodesToExplore = nodesToExplore[1:]
		i, j := currNode[0], currNode[1]
		if visited[i][j] {
			continue
		}
		visited[i][j] = true
		if matrix[i][j] == 0 {
			continue
		}
		currSize += 1

		unvisitedNeighbors := getUnvisited(i, j, matrix, visited)
		for _, neighbor := range unvisitedNeighbors {
			nodesToExplore = append(nodesToExplore, neighbor)
		}

	}
	fmt.Print(currSize)
	if currSize > 0 {
		result = append(result, currSize)
	}
	return result
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

// expected := []int{1, 2, 2, 2, 5}
// 	input := [][]int{
// 		{1, 0, 0, 1, 0},
// 		{1, 0, 1, 0, 0},
// 		{0, 0, 1, 0, 1},
// 		{1, 0, 1, 0, 1},
// 		{1, 0, 1, 1, 0},
// 	}
