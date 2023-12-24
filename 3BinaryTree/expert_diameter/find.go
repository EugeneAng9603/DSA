// find longest diameter

// O(n), O(h)
package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

type TreeInfo struct {
	diameter int
	height   int
}

func BinaryTreeDiameter(tree *BinaryTree) int {
	// Write your code here.
	return getTreeInfo(tree).diameter
}

func getTreeInfo(tree *BinaryTree) TreeInfo {
	if tree == nil {
		return TreeInfo{}
	}

	leftInfo := getTreeInfo(tree.Left)
	rightInfo := getTreeInfo(tree.Right)

	longestPathThroughRoot := leftInfo.height + rightInfo.height
	currMaxDiameter := max(leftInfo.diameter, rightInfo.diameter)
	// decide to include current node or not
	currDiameter := max(longestPathThroughRoot, currMaxDiameter)
	currHeight := 1 + max(leftInfo.height, rightInfo.height)

	return TreeInfo{currDiameter, currHeight}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
