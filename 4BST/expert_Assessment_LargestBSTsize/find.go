// n, h

// to do it in O(N)
// for every recursion, check 4 things
// left,right Valid, currNode >= max of left, <= min or right
package main

import (
	"math"
)

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func LargestBSTSize(tree *BinaryTree) int {
	// Write your code here.
	return getTreeInfo(tree).maxBST
}

type treeInfo struct {
	isBST  bool
	min    int
	max    int
	maxBST int
}

func getTreeInfo(tree *BinaryTree) treeInfo {
	if tree == nil {
		return treeInfo{
			isBST:  true,
			min:    math.MaxInt32,
			max:    math.MinInt32,
			maxBST: 0,
		}
	}

	leftInfo := getTreeInfo(tree.Left)
	rightInfo := getTreeInfo(tree.Right)

	checkBst := leftInfo.isBST && rightInfo.isBST &&
		tree.Value > leftInfo.max && tree.Value <= rightInfo.min

	// update min, max and largestBST based on the info
	maxValue := findMax(tree.Value, leftInfo.max, rightInfo.max)
	minValue := findMin(tree.Value, leftInfo.min, rightInfo.min)

	currMax := 0
	if checkBst {
		currMax = 1 + leftInfo.maxBST + rightInfo.maxBST
	} else {
		currMax = findMax(leftInfo.maxBST, rightInfo.maxBST)
	}

	return treeInfo{
		isBST:  checkBst,
		max:    maxValue,
		min:    minValue,
		maxBST: currMax,
	}
}

func findMax(a int, others ...int) int {
	max := a
	for _, num := range others {
		if num > max {
			max = num
		}
	}
	return max
}

func findMin(a int, others ...int) int {
	min := a
	for _, num := range others {
		if num < min {
			min = num
		}
	}
	return min
}

// root := NewBinaryTree(20)
// 	root.Left = NewBinaryTree(7)
// 	root.Left.Left = NewBinaryTree(0)
// 	root.Left.Right = NewBinaryTree(8)
// 	root.Left.Right.Left = NewBinaryTree(7)
// 	root.Left.Right.Right = NewBinaryTree(9)
// 	root.Right = NewBinaryTree(10)
// 	root.Right.Left = NewBinaryTree(5)
// 	root.Right.Left.Left = NewBinaryTree(2)
// 	root.Right.Left.Left.Left = NewBinaryTree(1)
// 	root.Right.Left.Right = NewBinaryTree(5)
// 	root.Right.Right = NewBinaryTree(15)
// 	root.Right.Right.Left = NewBinaryTree(13)
// 	root.Right.Right.Left.Right = NewBinaryTree(14)
// 	root.Right.Right.Right = NewBinaryTree(22)

// 	actual := LargestBSTSize(root)
// 	require.Equal(t, 9, actual)
// root 10 is the largest BST with 9 nodes

//                                1
//               3                                         10
//     -8                  0                    5                   15
// -9      -7      0-2                      6     5-2         13         22
//                                      1-2                      14
