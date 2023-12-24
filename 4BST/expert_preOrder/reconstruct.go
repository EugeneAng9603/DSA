// Given preorder values, reconstruct a BST

// For BST: the right node is the first next larger value
// the next (not larger) value is left
// if we always find the next larger and next not larger,
// it turns out to be n^2, not efficient (see solution1)

// So, use min, max and recursion.; for every node
// its right has a min of itself , its left has a max itself
// we just check the arr[rootIdx] if it's valid in left or right
// else return nil and find another potential node
// RMB to keep track rootIdx too
// lastly, create a BST node according to the left, right received

package main

import (
	"math"
)

// This is an input class. Do not edit.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

type treeInfo struct {
	rootIdx int
}

func ReconstructBst(preOrderTraversalValues []int) *BST {
	// Write your code here.
	treeInfo := &treeInfo{0}

	return helper(
		math.MinInt32,
		math.MaxInt32,
		preOrderTraversalValues,
		treeInfo,
	)
}

func helper(min, max int, arr []int, info *treeInfo) *BST {
	// means already added all values in arr
	if info.rootIdx == len(arr) {
		return nil
	}

	rootValue := arr[info.rootIdx]
	// if out of bound, means cannot be under currnode
	if rootValue < min || rootValue >= max {
		return nil
	}

	info.rootIdx += 1
	// check the next arr[rootIdx] can insert into left or right
	// if cannot means left and right = nil
	left := helper(min, rootValue, arr, info)
	right := helper(rootValue, max, arr, info)

	return &BST{Value: rootValue, Left: left, Right: right}
}

// test case
// func getDfsOrder(node *BST, values []int) []int {
// 	if node == nil {
// 		return nil
// 	}
// 	values = append(values, node.Value)
// 	getDfsOrder(node.Left, values)
// 	getDfsOrder(node.Right, values)
// 	return values
// }

// func (s *TestSuite) TestCase1(t *TestCase) {
// 	preOrderTraversalValues := []int{10, 4, 2, 1, 3, 17, 19, 18}
// 	tree := &BST{Value: 10}
// 	tree.Left = &BST{Value: 4}
// 	tree.Left.Left = &BST{Value: 2}
// 	tree.Left.Left.Left = &BST{Value: 1}
// 	tree.Left.Right = &BST{Value: 3}
// 	tree.Right = &BST{Value: 17}
// 	tree.Right.Right = &BST{Value: 19}
// 	tree.Right.Right.Left = &BST{Value: 18}
// 	expected := getDfsOrder(tree, nil)
// 	actual := ReconstructBst(preOrderTraversalValues)
// 	actualOrder := getDfsOrder(actual, nil)
// 	require.Equal(t, expected, actualOrder)
// }
