// Recursively

package main

import "fmt"

// Do not edit the class below except for
// the insert, contains, and remove methods.
// Feel free to add new properties and methods
// to the class.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	// Write your code here.
	// Do not edit the return statement of this method.
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

func (tree *BST) Contains(value int) bool {
	// Write your code here.
	if value < tree.Value {
		if tree.Left == nil {
			return false
		} else {
			return tree.Left.Contains(value)
		}
	} else if value > tree.Value {
		if tree.Right == nil {
			return false
		} else {
			return tree.Right.Contains(value)
		}
	}
	return true
}

func (tree *BST) Remove(value int) *BST {
	// Write your code here.
	// Do not edit the return statement of this method.
	tree.removeWithParent(value, nil)
	return tree
}

// binary search the value
// if root to be remove
// if found value is on the Left / Right
func (tree *BST) removeWithParent(value int, parent *BST) {
	if value < tree.Value {
		if tree.Left != nil {
			tree.Left.removeWithParent(value, tree)
		}
	} else if value > tree.Value {
		if tree.Right != nil {
			tree.Right.removeWithParent(value, tree)
		}
	} else { // found! try to replace it with smallest val on right
		// case 1: with 2 nodes
		if tree.Left != nil && tree.Right != nil {
			tree.Value = tree.Right.getMinValue()
			tree.Right.removeWithParent(tree.Value, tree)
			// case 2: if remove root node with only 1 child
		} else if parent == nil {
			if tree.Left != nil {
				tree.Value = tree.Left.Value
				tree.Right = tree.Left.Right
				tree.Left = tree.Left.Left
			} else if tree.Right != nil { // if left is empty
				tree.Value = tree.Right.Value
				tree.Left = tree.Right.Left
				tree.Right = tree.Right.Right
			} else {
				// single-noed tree, do nth
			}
			// case 2: on the left and with 1 child only
		} else if parent.Left == tree {
			if tree.Left != nil {
				parent.Left = tree.Left
			} else {
				parent.Left = tree.Right
			}
			// case 3: on the right and with 1 child only
		} else if parent.Right == tree {
			if tree.Right != nil {
				parent.Right = tree.Left
			} else {
				parent.Right = tree.Right
			}
		}
	}
}

func (tree *BST) getMinValue() int {
	if tree.Left == nil {
		return tree.Value
	}
	return tree.Left.getMinValue()
}

func NewBST(value int) *BST {
	return &BST{Value: value}
}

func main() {
	root := NewBST(10)
	root.Left = NewBST(5)
	root.Left.Left = NewBST(2)
	root.Left.Left.Left = NewBST(1)
	root.Left.Right = NewBST(5)
	root.Right = NewBST(15)
	root.Right.Left = NewBST(13)
	root.Right.Left.Right = NewBST(14)
	root.Right.Right = NewBST(22)

	root.Insert(12)
	fmt.Print(root.Value, root.Left.Left.Value, root.Right.Value)
	fmt.Print("\n")
	fmt.Print(root.Left.Left.Value, root.Left.Right.Value, root.Right.Left.Value, root.Right.Right.Value)
	fmt.Print("\n")

	root.Remove(10)
	fmt.Print(root, root.Left, root.Right)
	fmt.Print("\n")
	fmt.Print(root.Left.Left, root.Left.Right, root.Right.Left, root.Right.Right)
	fmt.Print("\n")

	root.Contains(15)
	fmt.Print(root, root.Left, root.Right)
	fmt.Print("\n")
	fmt.Print(root.Left.Left, root.Left.Right, root.Right.Left, root.Right.Right)
	fmt.Print("\n")
}
