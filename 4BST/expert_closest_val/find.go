package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	// Write your code here.
	return tree.helper(target, tree.Value)
}

func (tree *BST) helper(target, closest int) int {
	if absdiff(target, closest) > absdiff(target, tree.Value) {
		closest = tree.Value
	}
	if target < tree.Value && tree.Left != nil {
		return tree.Left.helper(target, closest)
	} else if target > tree.Value && tree.Right != nil {
		return tree.Right.helper(target, closest)
	}
	return closest
}

func absdiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
