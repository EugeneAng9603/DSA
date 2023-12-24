// find max path sum

package main

import (
	"fmt"
	"math"
)

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func MaxPathSum(tree *BinaryTree) int {
	// Write your code here.
	_, maxSum := findMaxSum(tree)
	return maxSum
}

func findMaxSum(tree *BinaryTree) (int, int) {
	fmt.Print(tree)
	if tree == nil {
		return 0, math.MinInt32
	}

	leftMaxBranch, leftMaxPath := findMaxSum(tree.Left)
	rightMaxBranch, rightMaxPath := findMaxSum(tree.Right)
	maxChildSum := max(leftMaxBranch, rightMaxBranch)

	value := tree.Value
	maxSumAsBranch := max(maxChildSum+value, value)
	maxSumAsRootNode := max(leftMaxBranch+value+rightMaxBranch, maxSumAsBranch)
	maxPathSum := max(leftMaxPath, rightMaxPath, maxSumAsRootNode)

	return maxSumAsBranch, maxPathSum
}

func max(first int, vals ...int) int {
	for _, val := range vals {
		if val > first {
			first = val
		}
	}
	return first
}
