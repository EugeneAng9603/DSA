// traverse graph but cannot cross the diagonal
// we can find the half part first, then multiply by 2
package main

import "fmt"

func NumberOfWaysNoCrossDiagonal(n int) int {
	// Write your code here.
	// table := make([]int, n+1)

	return dp(n, n) * 2
}

// we try do left bottom half
func dp(row, col int) int {
	if row == 0 && col == 0 {
		return 1
	}
	if row < 0 || col < 0 {
		return 0
	}
	if row < col {
		return 0
	}
	return dp(row-1, col) + dp(row, col-1)
}
func main() {
	fmt.Print(NumberOfWaysNoCrossDiagonal(3)) // 10
	fmt.Print("\n")
	fmt.Print(NumberOfWaysNoCrossDiagonal(1)) // 2
	fmt.Print("\n")
	fmt.Print(NumberOfWaysNoCrossDiagonal(12)) // 616024
}
