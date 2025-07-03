// find order of biggest plus sign in a 2D matrix
// ✅ Time & Space Complexity
// Time: O(n × m)

// Space: O(n × m) → reduced from 4 × O(n × m) to 1 × O(n × m)
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

// ✅ Time & Space Complexity
// Time: O(n × m) — One pass for each of the 4 directions.

// Space: O(n × m) — For the 4 direction matrices.

// exceeded time limit
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func largestPlusSign(grid [][]int) int {
// 	n := len(grid)
// 	if n == 0 || len(grid[0]) == 0 {
// 		return 0
// 	}
// 	m := len(grid[0])

// 	// Initialize DP arrays
// 	left := make([][]int, n)
// 	right := make([][]int, n)
// 	up := make([][]int, n)
// 	down := make([][]int, n)

// 	for i := 0; i < n; i++ {
// 		left[i] = make([]int, m)
// 		right[i] = make([]int, m)
// 		up[i] = make([]int, m)
// 		down[i] = make([]int, m)
// 	}

// 	// Fill left and up
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			if grid[i][j] == 1 {
// 				left[i][j] = 1
// 				up[i][j] = 1
// 				if j > 0 {
// 					left[i][j] += left[i][j-1]
// 				}
// 				if i > 0 {
// 					up[i][j] += up[i-1][j]
// 				}
// 			}
// 		}
// 	}

// 	// Fill right and down
// 	for i := n - 1; i >= 0; i-- {
// 		for j := m - 1; j >= 0; j-- {
// 			if grid[i][j] == 1 {
// 				right[i][j] = 1
// 				down[i][j] = 1
// 				if j < m-1 {
// 					right[i][j] += right[i][j+1]
// 				}
// 				if i < n-1 {
// 					down[i][j] += down[i+1][j]
// 				}
// 			}
// 		}
// 	}

// 	// Find max plus sign order
// 	maxOrder := 0
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			if grid[i][j] == 1 {
// 				order := min4(left[i][j], right[i][j], up[i][j], down[i][j])
// 				if order > maxOrder {
// 					maxOrder = order
// 				}
// 			}
// 		}
// 	}

// 	// Return order - 1 if you need arm length only, or full order including center
// 	return maxOrder
// }

// func min4(a, b, c, d int) int {
// 	return int(math.Min(float64(a), math.Min(float64(b), math.Min(float64(c), float64(d)))))
// }

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
		{0, 1, 1, 1, 1},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 0, 1},
		{0, 1, 0, 1, 0},
	}

	result := largestPlusSign(grid)
	fmt.Println("Largest Plus Sign Order:", result) // Output: Largest Plus Sign Order: 2
}
