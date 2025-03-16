// h+k, h , k is input param

package main

import "fmt"

// This is an input class. Do not edit.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

type treeInfo struct {
	numberOfVisited int
	lastVisitedNode *BST
}

func FindKthLargestValueInBst(tree *BST, k int) int {
	// Write your code here.
	treeInfo := treeInfo{0, tree}
	reverseInOrder(tree, k, &treeInfo)
	return treeInfo.lastVisitedNode.Value
}

func reverseInOrder(node *BST, k int, treeInfo *treeInfo) {
	if node == nil || treeInfo.numberOfVisited >= k {
		return
	}

	reverseInOrder(node.Right, k, treeInfo)
	if treeInfo.numberOfVisited < k {
		treeInfo.numberOfVisited += 1
		treeInfo.lastVisitedNode = node
		// reverseInOrder(node.Left, k, treeInfo)
		// this works too
	}
	reverseInOrder(node.Left, k, treeInfo)
}

func main() {
	root := &BST{Value: 15}
	root.Left = &BST{Value: 5}
	root.Left.Left = &BST{Value: 2}
	root.Left.Left.Left = &BST{Value: 1}
	root.Left.Left.Right = &BST{Value: 3}
	root.Left.Right = &BST{Value: 5}
	root.Right = &BST{Value: 20}
	root.Right.Left = &BST{Value: 17}
	root.Right.Right = &BST{Value: 22}
	k := 3
	// 	expected := 17
	actual := FindKthLargestValueInBst(root, k)
	// require.Equal(t, expected, actual)
	fmt.Print(actual)
}
