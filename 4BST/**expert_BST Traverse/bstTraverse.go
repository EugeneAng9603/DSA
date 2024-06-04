package main

import (
	"fmt"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) InOrderTraverse(array []int) []int {
	// Write your code here.
	if tree.Left != nil {
		array = tree.Left.InOrderTraverse(array)
	}
	array = append(array, tree.Value)
	if tree.Right != nil {
		array = tree.Right.InOrderTraverse(array)
	}
	return array
}

func (tree *BST) PreOrderTraverse(array []int) []int {
	// Write your code here.
	array = append(array, tree.Value)
	if tree.Left != nil {
		array = tree.Left.PreOrderTraverse(array)
	}
	if tree.Right != nil {
		array = tree.Right.PreOrderTraverse(array)
	}
	return array
}

func (tree *BST) PostOrderTraverse(array []int) []int {
	// Write your code here.
	if tree.Left != nil {
		array = tree.Left.PostOrderTraverse(array)
	}
	if tree.Right != nil {
		array = tree.Right.PostOrderTraverse(array)
	}
	array = append(array, tree.Value)
	return array
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
	root.Right.Right = NewBST(22)

	// inOrder := []int{1, 2, 5, 5, 10, 15, 22}
	// preOrder := []int{10, 5, 2, 1, 5, 15, 22}
	// postOrder := []int{1, 2, 5, 5, 22, 15, 10}
	fmt.Print(root.InOrderTraverse([]int{}))
	fmt.Print(root.PreOrderTraverse([]int{}))
	fmt.Print(root.PostOrderTraverse([]int{}))
}

// // If dont use receiver for inorder
// func inorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}

// 	left := inorderTraversal(root.Left)
// 	right := inorderTraversal(root.Right)

// 	output := make([]int, 0)

// 	output = append(output, left...)
// 	output = append(output, root.Val)
// 	output = append(output, right...)
// 	return output

// }

// // function scope for inorder
// func inorderTraversal(root *TreeNode) []int {
//     res := []int{}
//     var inorder func(root *TreeNode)
//     inorder = func(root *TreeNode) {
//         if root == nil { return }
//         inorder(root.Left)
//         res = append(res, root.Val)
//         inorder(root.Right)
//     }
//     inorder(root)
//     return res
// }

// // print dfs without receiver
// func DepthFirstSearch(root *BinaryTree) []int {
// 	res := []int{}
// 	var dfs func(node *BinaryTree)
// 	dfs = func(node *BinaryTree) {
// 		res = append(res, node.Value)
// 		if node.Left != nil {
// 			dfs(node.Left)
// 		}
// 		if node.Right != nil {
// 			dfs(node.Right)
// 		}
// 	}
// 	dfs(root)
// 	return res
// }

// // ----------------------------------------------------------------------
// // Test case for traversal DFS using Receiver
// type BinaryTree struct {
// 	Value       int
// 	Left, Right *BinaryTree
// }

// func (n *BinaryTree) DepthFirstSearch(array []int) []int {
// 	// Write your code here.
// 	array = append(array, n.Value)
// 	if n.Left != nil {
// 		array = n.Left.DepthFirstSearch(array)
// 	}
// 	if n.Right != nil {
// 		array = n.Right.DepthFirstSearch(array)
// 	}
// 	return array
// }

// ------------------------------------------------------------
// TEST CASE
// func main() {

// 	root := &BinaryTree{Value: 1}
// 	root.Left = &BinaryTree{Value: 2}
// 	root.Left.Left = &BinaryTree{Value: 4}
// 	root.Left.Left.Left = &BinaryTree{Value: 8}
// 	root.Left.Left.Right = &BinaryTree{Value: 9}
// 	root.Left.Right = &BinaryTree{Value: 5}
// 	root.Right = &BinaryTree{Value: 3}
// 	root.Right.Left = &BinaryTree{Value: 6}
// 	root.Right.Right = &BinaryTree{Value: 7}
// 	fmt.Print(root.DepthFirstSearch([]int{}))
//     fmt.Print(DepthFirstSearch(root))
// }
