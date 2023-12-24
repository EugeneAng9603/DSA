package main

import "fmt"

type TreeNodeClosures struct {
	Val   int
	Right *TreeNodeClosures
	Left  *TreeNodeClosures
}

func rightSideViewClosures(root *TreeNodeClosures) []int {
	levels := []int{}

	var walk func(n *TreeNodeClosures, depth int)
	walk = func(n *TreeNodeClosures, d int) {
		if n == nil {
			return
		}
		// If new level, add node to it
		// Since we traverse right node then left node, right gets inserted
		// first. And if there's no right node, left gets to go in.
		if len(levels) < d+1 {
			levels = append(levels, n.Val)
		}
		walk(n.Right, d+1)
		walk(n.Left, d+1)
		return
	}

	walk(root, 0)
	return levels
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
	node := &TreeNodeClosures{Val: 3}
	node.Right = &TreeNodeClosures{Val: 5}
	node.Left = &TreeNodeClosures{Val: 2}

	fmt.Print(rightSideViewClosures(node)) // Output: [1 2]

}
