// given parent, left, right, find successor

// LeftMostChild of the right node or RightMostParent of itself
// means if a node is in the right of its parent,
// keep searching parent until the node is not in the right of its parent
// then next should be node.parent
package main

// O(h), O(1)
// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree
}

func FindSuccessor(tree *BinaryTree, node *BinaryTree) *BinaryTree {
	// Write your code here.
	if node.Right != nil {
		return getLeftMostChild(node.Right)
	}
	return getRightMostParent(node)
}

func getLeftMostChild(node *BinaryTree) *BinaryTree {
	curr := node
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

// if node is on the right of its parent, keep searching next parent
// if on the left of its parent, just return node.parent
func getRightMostParent(node *BinaryTree) *BinaryTree {
	currNode := node
	for currNode.Parent != nil && currNode.Parent.Right == currNode {
		currNode = currNode.Parent
	}

	return currNode.Parent
}
