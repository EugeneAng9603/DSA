// return [15,16,18,10,11]

package main

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	// Write your code here.
	array := []int{}
	helper(root, 0, &array)
	return array
}

func helper(node *BinaryTree, runningSum int, array *[]int) {
	if node == nil {
		return
	}

	new_temp_sum := runningSum + node.Value
	if node.Left == nil && node.Right == nil {
		*array = append(*array, new_temp_sum)
	}
	helper(node.Left, new_temp_sum, array)
	helper(node.Right, new_temp_sum, array)
}
