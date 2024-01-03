// O(h), O(1)
// given BST and three node, check if they are 1->2->3 decendants or 3->2->1

// Normal way is to recursively or iteratively find descendant
// if one->two, then check two->three, else the other way
// Recursive: O(h), O(h), Iterative: O(h), O(1)
// Best way is to stop if one->three without seeing two
// So Time:O(d), distance between one and three
// See solution 3

package main

// This is an input class. Do not edit.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func ValidateThreeNodes(nodeOne *BST, nodeTwo *BST, nodeThree *BST) bool {
	// Write your code here.
	if isDes(nodeOne, nodeTwo) {
		return isDes(nodeTwo, nodeThree)
	}
	if isDes(nodeTwo, nodeOne) {
		return isDes(nodeThree, nodeTwo)
	}
	return false
}

// Iterative
func isDes(node, target *BST) bool {
	currNode := node
	for currNode != nil && currNode != target {
		if target.Value < currNode.Value {
			currNode = currNode.Left
		} else {
			currNode = currNode.Right
		}
	}
	return currNode == target
}

// root := &BST{Value: 5}
// 	root.Left = &BST{Value: 2}
// 	root.Right = &BST{Value: 7}
// 	root.Left.Left = &BST{Value: 1}
// 	root.Left.Right = &BST{Value: 4}
// 	root.Right.Left = &BST{Value: 6}
// 	root.Right.Right = &BST{Value: 8}
// 	root.Left.Left.Left = &BST{Value: 0}
// 	root.Left.Right.Left = &BST{Value: 3}

// 	nodeOne := root
// 	nodeTwo := root.Left
// 	nodeThree := root.Left.Right.Left
// 	expected := true
