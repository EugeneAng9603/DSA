// You are given an array of arrays.
// Each subarray in this array holds two integer values and represents an item;
// the rst integer is the item's value, and the second integer is the item's weight.
// You are also given an integer representing the maximum capacity of a knapsack that you have.
// Your goal is to t items in your knapsack, all the while maximizing their combined value.
// Note that the sum of the weights of the items that you pick cannot exceed the knapsack's capacity.
// Write a function that returns the maximized combined value of the items that you should pick,
// as well as an array of the indices of each item picked. Assume that there will only
// be one combination of items that maximizes the total value in the knapsack.

// Sample input: [[1, 2], [4, 3], [5, 6], [6, 7]], 10 Sample output: [10, [1, 3]]

// O(nc), O(nc), n is number of items, c is capacity

package main

import "fmt"

//    		0   1  2  3  4  5  6  7  8  9 10
//[]		0   x  x  x  x  x  x  x  x  x  x
//[1,2]		1   1  1  1  x  x  x  x  x  x  x
//[4,3]
//[5,6]
//[6,7]

// this is a classic 0/1 knapsack problem solved using dynamic programming
func FindKnapsackMaxValue(items [][]int, capacity int) (int, []int) {
	n := len(items)
	if n == 0 || capacity <= 0 {
		return 0, []int{}
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		value, weight := items[i-1][0], items[i-1][1]
		for w := 0; w <= capacity; w++ {
			if weight > w {
				dp[i][w] = dp[i-1][w]
			} else {
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-weight]+value)
			}
		}
	}

	maxValue := dp[n][capacity]
	selectedItems := []int{}
	w := capacity

	for i := n; i > 0 && maxValue > 0; i-- {
		if maxValue != dp[i-1][w] {
			selectedItems = append(selectedItems, i-1)
			maxValue -= items[i-1][0]
			w -= items[i-1][1]
		}
	}

	return dp[n][capacity], reverseSlice(selectedItems)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverseSlice(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
func main() {
	items := [][]int{{1, 2}, {4, 3}, {5, 6}, {6, 7}}
	capacity := 10
	maxValue, selectedItems := FindKnapsackMaxValue(items, capacity)
	println("Max Value:", maxValue)
	fmt.Printf("Selected Items Indices:%v", selectedItems)
}
