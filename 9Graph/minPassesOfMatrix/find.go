// O(w*h)
// we loop through all to find positive and add to queue
// then for each in the queue, keep adding to the queue
// for each positive, check its neighbors and convert those -ve to +ve
// once done first round, passes +1 then proceed to second passes
// do this by using a currQueueSize to differentiate queue within 1 pass

// Bonus: to use just one queue, change condition to for i < len(size)
// so even added new element in the queue, we will start next for loop
// for those passed in the next loop

// on every pass, for every non 0 +ve node, we can convert its neighbors to +ve
// once converted, we dont use it to convert its neighbors in a same pass
//
//	input := [][]int{
//		{0, -1, -3, 2, 0},
//		{1, -2, -5, -1, -3},
//		{3, 0, 0, -4, -1},
//	}
//
// expected := 3
package main

import (
	"fmt"
)

type intPair struct {
	row, col int
}

func MinimumPassesOfMatrix(matrix [][]int) int {
	// Write your code here.
	passes := convertNegative(matrix)

	// if no more -ve values
	if !containsNegative(matrix) {
		// bcuz we loop through elements in last queue, which should do nth
		// all possible ones are converted to positive
		// so minus the last pass
		return passes - 1
	} else {
		return -1
	}
}

func convertNegative(matrix [][]int) int {
	queue := getAllPositive(matrix)

	var passes int
	for len(queue) > 0 {
		// keep track of current queue size to differentiate element
		// from this loop and next
		fmt.Print(queue)
		queueSize := len(queue)
		// for all +ve in 1 pass
		for queueSize > 0 {
			curr := queue[0]
			queue = queue[1:]
			currRow, currCol := curr.row, curr.col

			// get neighbors
			adjacents := getNeighbors(matrix, currRow, currCol)

			for _, position := range adjacents {
				// check if it's negative,
				// if -ve , change to positive, append to queue
				i, j := position.row, position.col
				if matrix[i][j] < 0 {
					matrix[i][j] *= -1
					queue = append(queue, position)
				}
			}
			// rmb to decrement so if queueSize is 0,
			// start next loop
			queueSize -= 1
		}
		passes += 1
	}
	return passes
}

func getNeighbors(matrix [][]int, row, col int) []intPair {
	neighbors := make([]intPair, 0)
	rowlen := len(matrix)
	collen := len(matrix[row])

	if row > 0 {
		neighbors = append(neighbors, intPair{row - 1, col}) // UP
	}
	if col > 0 {
		neighbors = append(neighbors, intPair{row, col - 1}) // LEFT
	}
	if row < rowlen-1 {
		neighbors = append(neighbors, intPair{row + 1, col}) // DOWN
	}
	if col < collen-1 {
		neighbors = append(neighbors, intPair{row, col + 1}) // RIGHT
	}
	return neighbors
}

func getAllPositive(matrix [][]int) []intPair {
	pairs := []intPair{}
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] > 0 {
				pairs = append(pairs, intPair{row, col})
			}
		}
	}
	return pairs
}

func containsNegative(matrix [][]int) bool {
	for _, row := range matrix {
		for _, value := range row {
			if value < 0 {
				return true
			}
		}
	}
	return false
}
