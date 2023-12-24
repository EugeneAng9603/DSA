// flatten, then connect with doubly linked list

// O(n), O(d)
// if use in-order and store in arr, uses O(n) space.
// if only store left most of right and rightmost of left, space = O(d)
package main

// This is the class of the input root. Do not edit it.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func FlattenBinaryTree(root *BinaryTree) *BinaryTree {
	// Write your code here.
	leftMost, _ := flatten(root)
	return leftMost
}

// return leftmost and rightmost for every recursion
// start from currNode, leftsub leftmost and right most
// then connect, then update leftMost (for return purpose)
// same for rightsub
func flatten(node *BinaryTree) (leftMost, rightMost *BinaryTree) {
	leftMost = node
	if node.Left != nil {
		leftSubLeftMost, leftSubRightMost := flatten(node.Left)
		connectNodes(leftSubRightMost, node)
		leftMost = leftSubLeftMost
	}

	rightMost = node
	if node.Right != nil {
		rightSubLeftMost, rightSubRightMost := flatten(node.Right)
		connectNodes(node, rightSubLeftMost)
		rightMost = rightSubRightMost
	}
	return leftMost, rightMost
}

func connectNodes(node1, node2 *BinaryTree) {
	node1.Right = node2
	node2.Left = node1
}
