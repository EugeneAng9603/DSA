// print all path of Btree
package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	output := make([]string, 0)

	binaryTreePathsUtil(root, "", &output)
	return output
}

func binaryTreePathsUtil(root *TreeNode, curr string, output *[]string) {
	if root == nil {
		return
	}

	valString := strconv.Itoa(root.Val)
	if curr == "" {
		curr = valString
	} else {
		curr = curr + "->" + valString
	}
	if root.Left == nil && root.Right == nil {
		*output = append(*output, curr)
		return
	}

	binaryTreePathsUtil(root.Left, curr, output)
	binaryTreePathsUtil(root.Right, curr, output)

}

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	output := binaryTreePaths(&root)
	fmt.Println(output)
}
