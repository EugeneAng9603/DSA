// Time complexity: O(N!)
// Space complexity: O(N^2)
package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	// Create an empty board of size n x n filled with dots
	board := make([]string, n)
	for i := range board {
		board[i] = strings.Repeat(".", n)
	}
	var result [][]string
	backtrack(0, &result, board, n)
	return result
}

func backtrack(row int, result *[][]string, board []string, n int) {
	// If we reached the end of the board, we found a valid solution
	if row == n {
		// Create a deep copy of the current board state
		temp := make([]string, n)
		copy(temp, board)
		*result = append(*result, temp)
		return
	}

	// Try placing a queen in each column of the current row
	for col := 0; col < n; col++ {
		if isValid(board, row, col, n) {
			// Convert string to byte slice to modify single character
			newRow := []byte(board[row])
			// Place the queen
			newRow[col] = 'Q'
			board[row] = string(newRow)

			// Move to next row
			backtrack(row+1, result, board, n)

			// Backtrack: remove the queen
			newRow[col] = '.'
			board[row] = string(newRow)
		}
	}
}

// isValid checks if placing a queen at given position is safe
func isValid(board []string, row, col, n int) bool {
	// Check all rows above in the same column
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// Check upper-left diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// Check upper-right diagonal
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// Position is safe if no conflicts found
	return true
}

func main() {
	fmt.Println(solveNQueens(5))
}
