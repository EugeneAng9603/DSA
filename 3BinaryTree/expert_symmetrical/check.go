// both O(n), O(h)

// Recursive
// left == right and left.left == right.right, left.right == right.left
package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func SymmetricalTree(tree *BinaryTree) bool {
	// Write your code here.
	return helper(tree.Left, tree.Right)
}

func helper(left, right *BinaryTree) bool {
	if left != nil && right != nil && left.Value == right.Value {
		return helper(left.Left, right.Right) && helper(left.Right, right.Left)
	}
	return left == right
}

// ------------------------------------------------------------------------------------------------------------
// Iteration, queue

// package main

// // This is an input class. Do not edit.
// type BinaryTree struct {
// 	Value int

// 	Left  *BinaryTree
// 	Right *BinaryTree
// }

// func SymmetricalTree(tree *BinaryTree) bool {
// 	// Write your code here.
//     stack1 := []*BinaryTree{tree.Left}
//     stack2 := []*BinaryTree{tree.Right}

//     for len(stack1) > 0 {
//         var left, right *BinaryTree
//         left, stack1 = stack1[len(stack1)-1], stack1[:len(stack1)-1]
//         right, stack2 = stack2[len(stack2)-1], stack2[:len(stack2)-1]

//         if left == nil && right == nil {
//             continue
//         }

//         if left == nil || right == nil || left.Value != right.Value {
//             return false
//         }

//         stack1 = append(stack1, left.Left, left.Right)
//         stack2 = append(stack2, right.Right, right.Left)
//     }
// 	return true
// }
