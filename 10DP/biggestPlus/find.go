// find order of biggest plus sign in a 2D matrix
package main

import (
	"fmt"
	"math"
)

func largestPlusSign(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])

	// Step 1: Initialize result matrix with max values
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				dp[i][j] = math.MaxInt32 // will be reduced to min of 4 directions
			}
		}
	}

	// Left pass
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				count++
				dp[i][j] = min(dp[i][j], count)
			} else {
				count = 0
			}
		}
	}

	// Right pass
	for i := 0; i < n; i++ {
		count := 0
		for j := m - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				count++
				dp[i][j] = min(dp[i][j], count)
			} else {
				count = 0
			}
		}
	}

	// Up pass
	for j := 0; j < m; j++ {
		count := 0
		for i := 0; i < n; i++ {
			if grid[i][j] == 1 {
				count++
				dp[i][j] = min(dp[i][j], count)
			} else {
				count = 0
			}
		}
	}

	// Down pass
	maxOrder := 0
	for j := 0; j < m; j++ {
		count := 0
		for i := n - 1; i >= 0; i-- {
			if grid[i][j] == 1 {
				count++
				dp[i][j] = min(dp[i][j], count)
				if dp[i][j] > maxOrder {
					maxOrder = dp[i][j]
				}
			} else {
				count = 0
			}
		}
	}

	return maxOrder
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// grid := [][]int{
	// 	{1, 0, 1, 0, 0},
	// 	{1, 1, 1, 1, 1},
	// 	{1, 0, 1, 0, 0},
	// 	{0, 1, 1, 1, 1},
	// 	{0, 0, 0, 1, 0},
	// }
	grid := [][]int{
		{0, 1, 1, 0, 1},
		{0, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 0, 1},
		{0, 1, 1, 1, 0},
	}

	result := largestPlusSign(grid)
	fmt.Println("Largest Plus Sign Order:", result) // Output: Largest Plus Sign Order: 2
}
