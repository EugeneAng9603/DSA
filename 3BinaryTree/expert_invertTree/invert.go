// mirror invert
// n, d
package main

import "fmt"

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

// func (tree *BinaryTree) InvertBinaryTree() {
// 	// Write your code here.
// 	tree.Left, tree.Right = tree.Right, tree.Left
// 	if tree.Left != nil {
// 		tree.Left.InvertBinaryTree()
// 	}
// 	if tree.Right != nil {
// 		tree.Right.InvertBinaryTree()
// 	}
// }

func (tree *BinaryTree) InvertBinaryTree() {
	if tree.Left != nil {
		tree.Left.InvertBinaryTree()
	}
	if tree.Right != nil {
		tree.Right.InvertBinaryTree()
	}
	tree.Left, tree.Right = tree.Right, tree.Left

}

func (tree *BinaryTree) PrintElement() {
	if tree == nil {
		return
	}
	queue := []*BinaryTree{tree}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node.Value, ", ")
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	// fmt.Println()
}

func main() {
	node := &BinaryTree{Value: 3}
	node.Right = &BinaryTree{Value: 5}
	node.Left = &BinaryTree{Value: 1}
	node.Right.Left = &BinaryTree{Value: 4}
	node.Left.Right = &BinaryTree{Value: 2}
	node.Right.Left.Left = &BinaryTree{Value: 6}
	node.Right.Left.Right = &BinaryTree{Value: 10}

	node.InvertBinaryTree()
	node.PrintElement()
}

//             3
//     1               5
//        2        4
//                6 10
