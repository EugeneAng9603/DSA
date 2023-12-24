// Given arr, construct min height BST
package main

func MinHeightBST(array []int) *BST {
	// Write your code here.
	return constructMin(array, 0, len(array)-1)
}

func constructMin(array []int, startIdx, endIdx int) *BST {
	if startIdx > endIdx {
		return nil
	}

	currIdx := (startIdx + endIdx) / 2
	bst := &BST{Value: array[currIdx]}
	bst.Left = constructMin(array, startIdx, currIdx-1)
	bst.Right = constructMin(array, currIdx+1, endIdx)
	return bst
}

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}
