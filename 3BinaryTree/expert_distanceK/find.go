// if only find those below, see my solution 2
// if consider those from above or another sub-tree, need to use parentNode
// or keep track of distance
// Track visited, see if target in left or right subtree or not at all
// node==target, search subtree
// target in leftsub, look for node distance k-L-1 on right

// O(n), O(n)
package main

import (
	"fmt"
)

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func FindNodesDistanceK(tree *BinaryTree, target int, k int) []int {
	// Write your code here.
	result := []int{}
	findDistance(tree, target, k, &result)
	return result
}

// findDistance from current node to target node(leftDistance and rightDistance)
func findDistance(node *BinaryTree, target, k int, result *[]int) int {

	if node == nil {
		return -1
	}
	// find those under it
	if node.Value == target {
		addNodeAtDistanceK(node, 0, k, result)
		return 1
	}

	// !!!! distance of node to target, not node.left to target
	leftDistance := findDistance(node.Left, target, k, result)
	rightDistance := findDistance(node.Right, target, k, result)

	//4 cases

	// case 1 , either left or right reach
	if leftDistance == k || rightDistance == k {
		*result = append(*result, node.Value)
	}
	// case 2, target in left sub, not in right sub
	if leftDistance != -1 {
		addNodeAtDistanceK(node.Right, leftDistance+1, k, result)
		fmt.Print("in left", node.Value, leftDistance+1)
		return leftDistance + 1
	}
	// case 3, target in right sub, not in left sub
	if rightDistance != -1 {
		// e.g. root: rightdistance is 2 (to node 6), target is right
		// addNode(node.Left), distance is right + 1 = 3
		// cuz we must start addNode for node.Left,
		// distance to root from node6 is 2 , so distance to node.Left is 3
		fmt.Print("in right", node.Value, rightDistance+1)
		addNodeAtDistanceK(node.Left, rightDistance+1, k, result) // cuz we check node.Left, 1 closer to k
		return rightDistance + 1
	}

	return -1
}

// this function explore all nodes from "distance" away from node
// distance+1 when looking deep to find a kth distance ones
// if distance to node == k, e.g. distance = 2, append these values
// else keep finding
func addNodeAtDistanceK(node *BinaryTree, distance int, k int, nodesDistanceK *[]int) {
	if node == nil {
		return
	}
	if distance == k {
		*nodesDistanceK = append(*nodesDistanceK, node.Value)
	} else {
		addNodeAtDistanceK(node.Left, distance+1, k, nodesDistanceK)
		addNodeAtDistanceK(node.Right, distance+1, k, nodesDistanceK)
	}
}

//                 1     (right distance = 1 to node 3, target is right)
//      2                     3
//   4     5                        6
//                               7     8
//test case
// root := &BinaryTree{Value: 1}
// 	root.Left = &BinaryTree{Value: 2}
// 	root.Right = &BinaryTree{Value: 3}
// 	root.Left.Left = &BinaryTree{Value: 4}
// 	root.Left.Right = &BinaryTree{Value: 5}
// 	root.Right.Right = &BinaryTree{Value: 6}
// 	root.Right.Right.Left = &BinaryTree{Value: 7}
// 	root.Right.Right.Right = &BinaryTree{Value: 8}
// 	target := 3
// 	k := 2
// 	expected := []int{2, 7, 8}

// ------------------------------------------------------------------------------------------------------------
// solution 2 only to find those below
// if only consider those under target

// type BinaryTree struct {
// 	Value int

// 	Left  *BinaryTree
// 	Right *BinaryTree
// }

// func FindNodesDistanceK(tree *BinaryTree, target int, k int) []int {
// 	// Write your code here.
//     result := []int{}
//     helper(tree, target, k, &result, false)
// 	return result
// }

// func helper(tree *BinaryTree, target int, k int, result *[]int, found bool) {
//     if tree == nil || k < 0{
//         return
//     }
//     if k == 0 {
//         *result = append(*result, tree.Value)
//         return
//     }
//     if tree.Value == target {
//         found = true
//     }
//     if found {
//         helper(tree.Left, target, k-1, result, found)
//         helper(tree.Right, target, k-1, result, found)
//     } else {
//         helper(tree.Left, target, k, result, found)
//         helper(tree.Right, target, k, result, found)
//     }
//     return
// }
