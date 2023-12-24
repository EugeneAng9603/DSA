// sum of all node depth of each subtree
// level 1: 1 -> 2,3
// level 2: 2-> 4,5   3-> 6,7
// level 3: 4-> 8,9
// depth: 1 is 0, 2,3 is 1, 4,5,6,7 is 2, 8,9 is 3
// output is 26 = node 1 + node 2 + node 3 + node 4 (node 5,6,7,8,9 has no left and right)
//                  16   +   2    +    2   +   6    (cuz node 4 is 3 + 3, node 1 is 2+2+2+2+3+3)

// Solution 1: apply simple node depth recursion on every node
// simple node depth formula is: N(node, depth) = depth + N(node.Left, depth+1) + N(node.Right, depth+1)
// since we are applying this to every subtree, average Time: O(nlogn), worst O(n2) is not balanced, space: O(h)

// func AllKindsOfNodeDepths(root *BinaryTree) int {
// 	// Write your code here.
//     if root == nil {
//         return 0
//     }
// 	return AllKindsOfNodeDepths(root.Left) + AllKindsOfNodeDepths(root.Right) + helper(root, 0)
// }

// func helper(node *BinaryTree, depth int) int {
//     if node == nil {
//         return 0
//     }
//     return helper(node.Left, depth+1) + helper(node.Right, depth+1) + depth
// }

// ------------------------------------------------------------------------------------------------
// Solution 2:
// understand that level 1 = level 2 left + level 2 right
// the result of level 1 = level 2 left + number of nodes(of left + its sub nodes) + level 2 right + number of nodes
// because each level 2 is 1 level relative to root, level 3 is 2 level relative to root etc....
// N = N.Left + # of Left + N.Right + # of Right

// Implementation 1: O(n), O(h) calculate depth and # of nodes, then sum it up for each subtree
// three pass of the entire tree
// 1: calculate number of nodes involved of each node
// 2: calculate node depth of each subtree using the N formula
// 3: calculate sum of all depth of each subtree
package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

type treeInfo struct {
	numNodes       int
	sumOfDepths    int
	sumOfAllDepths int
}

func AllKindsOfNodeDepths(root *BinaryTree) int {
	return getTreeInfo(root).sumOfAllDepths
}

func getTreeInfo(tree *BinaryTree) treeInfo {
	if tree == nil {
		return treeInfo{}
	}

	leftInfo, rightInfo := getTreeInfo(tree.Left), getTreeInfo(tree.Right)

	sumOfLeftDepths := leftInfo.sumOfDepths + leftInfo.numNodes
	sumOfRightDepths := rightInfo.sumOfDepths + rightInfo.numNodes

	numNodesInTree := 1 + leftInfo.numNodes + rightInfo.numNodes
	sumOfDepths := sumOfLeftDepths + sumOfRightDepths
	sumOfAllDepths := sumOfDepths + leftInfo.sumOfAllDepths + rightInfo.sumOfAllDepths

	return treeInfo{numNodes: numNodesInTree, sumOfDepths: sumOfDepths, sumOfAllDepths: sumOfAllDepths}
}

// --------------------------------------------------------------------------------
// Simple Recursion by passing down depthSum
// O(n), O(h)
// func AllKindsOfNodeDepths(root *BinaryTree) int {
// 	return helper(root, 0, 0)
// }

// func helper(root *BinaryTree, depthSum, depth int) int {
//     if root == nil {
//         return 0
//     }

//     depthSum += depth
//     return depthSum + helper(root.Left, depthSum, depth+1) + helper(root.Right, depthSum, depth+1)
// }

// --------------------------------------------------------------------------------
// Math logic
// O(n), O(h)
// func AllKindsOfNodeDepths(root *BinaryTree) int {
// 	// Write your code here.
// 	return helper(root, 0)
// }

// func helper(root *BinaryTree, depth int) int {
//     if root == nil {
//         return 0
//     }

//     // formula for 1+2+3+...+depth-1+depth
//     depthSum := (depth * (depth+1)) / 2

//     return depthSum + helper(root.Left, depth+1) + helper(root.Right, depth+1)
// }
