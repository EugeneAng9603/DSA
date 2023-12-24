// remove one edge, if left and right sum equal, true

// main idea is to calculate if there is a node with sum = whole tree / 2
// if yes then true
// where Bottom-Up, each node will return a sum of itself and its nodes
package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func SplitBinaryTree(tree *BinaryTree) int {
	// Write your code here.
	sum := Sum(tree)
	// if odd sum, means no way to split into half
	if sum%2 != 0 {
		return 0
	}

	desiredSum := sum / 2
	_, can := check(tree, desiredSum)
	if can {
		return desiredSum
	}
	return 0
}

// check if can split e.g. single node cannot split
// check if the desiredSum can be found in any subtree
func check(node *BinaryTree, target int) (int, bool) {
	if node == nil {
		return 0, false
	}
	leftSum, leftCan := check(node.Left, target)
	rightSum, rightCan := check(node.Right, target)

	currSum := leftSum + node.Value + rightSum

	return currSum, currSum == target || leftCan || rightCan
}

func Sum(node *BinaryTree) int {
	if node == nil {
		return 0
	}
	return node.Value + Sum(node.Left) + Sum(node.Right)
}
