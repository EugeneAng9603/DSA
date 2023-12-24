package main

import "fmt"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func rightSideView(root *TreeNode) *[]int {
	result := &[]int{}
	helper(root, result, 0)
	return result
}

func helper(node *TreeNode, result *[]int, depth int) {
	if node == nil {
		return
	}
	if depth == len(*result) {
		*result = append(*result, node.Val)
	}
	// fmt.Print(result, depth)
	// fmt.Print("\n")
	helper(node.Right, result, depth+1)
	helper(node.Left, result, depth+1)
}

// func newTree(value int) *TreeNode {
// 	return &TreeNode{Val: value}
// }

func main() {
	// node := newTree(3)
	// node.Right = newTree(5)
	// node.Left = newTree(1)
	// node.Right.Left = newTree(4)
	// node.Left.Right = newTree(2)
	// node.Right.Left.Left = newTree(6)
	// node.Right.Left.Right = newTree(10)

	//Optional
	node := TreeNode{Val: 3}
	node.Right = &TreeNode{Val: 5}
	node.Left = &TreeNode{Val: 2}

	fmt.Print(rightSideView(&node)) // Output: [1 2]

}
