// level traverse binary tree from the bottom
// so output should be [[15,7],[9,20],[3]]
// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
// option 1: use template, and prepend new slice to 2d slice
// option 2: use template, then reverse result
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := [][]int{}
	compute(&result, root, 0)
	return result
}

func compute(result *[][]int, node *TreeNode, level int) {
	if node == nil {
		return
	}
	// every start of the level
	// *** Tricky part is we must prepend, not append
	// so when we are left subtree 9, 9 is added to [ [9] [] ]
	// then later we prepend to 9 so 15,7 will be added to the first, 9 will be in the second subarray
	if len(*result) == level {
		// must prepend not append
		//*result = append(*result, []int{})
		newArr := [][]int{}
		newArr = append(newArr, []int{})
		newArr = append(newArr, *result...)
		*result = newArr
		//*result = append([][]int{}, *result...)
		//fmt.Print(*result)
	}

	compute(result, node.Left, level+1)
	compute(result, node.Right, level+1)
	length := len(*result)
	// fmt.Print(length, node.Val)
	// fmt.Print("\n")
	(*result)[length-level-1] = append((*result)[length-level-1], node.Val)
}

func main() {
	node := &TreeNode{3, nil, nil}
	node.Left = &TreeNode{9, nil, nil}
	node.Right = &TreeNode{20, nil, nil}
	node.Right.Left = &TreeNode{15, nil, nil}
	node.Right.Right = &TreeNode{7, nil, nil}

	fmt.Print(levelOrderBottom(node))
}
