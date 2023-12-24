// Can use function without receiver too, SEE BELOW!

package main

import "math"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) ValidateBst() bool {
	// Write your code here.
	return tree.helper(math.MinInt32, math.MaxInt32)
}

func (tree *BST) helper(min, max int) bool {
	if tree.Value < min || tree.Value >= max {
		return false
	}
	if tree.Left != nil && !tree.Left.helper(min, tree.Value) {
		return false
	}
	if tree.Right != nil && !tree.Right.helper(tree.Value, max) {
		return false
	}
	return true
}

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type BST struct {
// 	Value int

// 	Left  *BST
// 	Right *BST
// }

// func NewBST(value int) *BST {
// 	return &BST{Value: value}
// }

// func ValidateBst(tree *BST) bool {
// 	// Write your code here.
// 	return helper(tree, math.MinInt32, math.MaxInt32)
// }

// func helper(tree *BST, min, max int) bool {
// 	if tree.Value < min || tree.Value >= max {
// 		return false
// 	}
// 	if tree.Left != nil && !helper(tree.Left, min, tree.Value) {
// 		return false
// 	}
// 	if tree.Right != nil && !helper(tree.Right, tree.Value, max) {
// 		return false
// 	}
// 	return true
// }

// func main() {
// 	root := NewBST(10)
// 	root.Left = NewBST(5)
// 	root.Left.Left = NewBST(2)
// 	root.Left.Left.Left = NewBST(1)
// 	root.Left.Right = NewBST(5)
// 	root.Right = NewBST(15)
// 	root.Right.Left = NewBST(13)
// 	root.Right.Left.Right = NewBST(14)
// 	root.Right.Right = NewBST(22)

// 	result := ValidateBst(root)
// 	fmt.Print(result)
// }
