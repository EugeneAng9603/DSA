// given tree, add all depth. root is 0, then next level is 1, next is 2

package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
	// Write your code here.
	return helper(root, 0)
}

func helper(node *BinaryTree, depth int) int {
	if node == nil {
		return 0
	}
	return depth + helper(node.Left, depth+1) + helper(node.Right, depth+1)
}
