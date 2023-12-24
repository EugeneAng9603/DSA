// -1 is add, -2 is minus etc...
// every node will be the operator, except bottom nodes
// bottom up to find the result

package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func EvaluateExpressionTree(tree *BinaryTree) int {
	// Write your code here.
	if tree.Value >= 0 {
		return tree.Value
	}

	leftValue := EvaluateExpressionTree(tree.Left)
	rightValue := EvaluateExpressionTree(tree.Right)

	if tree.Value == -1 {
		return leftValue + rightValue
	}
	if tree.Value == -2 {
		return leftValue - rightValue
	}
	if tree.Value == -3 {
		return leftValue / rightValue
	}
	return leftValue * rightValue
}
