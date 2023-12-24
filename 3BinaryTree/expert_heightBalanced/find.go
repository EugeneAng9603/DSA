// if balanced, means for all subtree, leftsub and rightsub height diff is <= 1

package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

type treeInfo struct {
	isBalanced bool
	height     int
}

func HeightBalancedBinaryTree(tree *BinaryTree) bool {
	// Write your code here.
	treeInfo := getTreeInfo(tree)
	return treeInfo.isBalanced
}

func getTreeInfo(tree *BinaryTree) treeInfo {
	if tree == nil {
		return treeInfo{
			isBalanced: true,
			height:     0,
		}
	}

	leftInfo := getTreeInfo(tree.Left)
	rightInfo := getTreeInfo(tree.Right)
	IsBalanced := leftInfo.isBalanced && rightInfo.isBalanced && abs(leftInfo.height-rightInfo.height) <= 1
	Height := max(leftInfo.height, rightInfo.height) + 1

	return treeInfo{
		isBalanced: IsBalanced,
		height:     Height,
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
