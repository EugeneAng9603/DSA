package main

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
		reverseInOrder(node.Left, k, treeInfo)
	}
}
