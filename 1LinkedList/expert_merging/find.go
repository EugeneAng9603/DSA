// find merging node
// 2,3,1,4 and 8,7,1,4
// output is 1,4

// Solution 1: use Set/Map to store ll1, if found same value from ll2, return
// Solution 2: Calculate the length of both, find diff, then get the longer
// one move to next nth first, then if node1 == node2 return
// Solution 3: Same as 2, but just get pointer1 move to ll2 to start when
// reached nil, same for pointer2. once meet, return

package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

// for the shorter one, reach nil first then jump to longer head
// then longer will reach nil and jump to shorter head
// now when they meet, that should be the merging one
func MergingLinkedLists(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
	// Write your code here.
	node1, node2 := linkedListOne, linkedListTwo

	for node2 != node1 {
		// if node1 shorter and reach nil, jump to head of ll2
		if node1 == nil {
			node1 = linkedListTwo
		} else {
			node1 = node1.Next
		}

		// if node2 shorter and reach nil, jump to head of ll1
		if node2 == nil {
			node2 = linkedListOne
		} else {
			node2 = node2.Next
		}
	}

	// when shorter jump to longer,
	return node1
}

// --------------------------------------------------------------------------------------------------------------------
// // This is an input struct. Do not edit.
// type LinkedList struct {
// 	Value int
// 	Next  *LinkedList
// }

// func MergingLinkedLists(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
// 	// Write your code here.
//     currOne, currTwo := linkedListOne, linkedListTwo
//     countOne, countTwo := 0, 0
//     for currOne != nil {
//         countOne += 1
//         currOne = currOne.Next
//     }
//     for currTwo != nil {
//         countTwo += 1
//         currTwo = currTwo.Next
//     }
//     diff := abs(countOne - countTwo)

//     currOne, currTwo = linkedListOne, linkedListTwo
//     var biggerNode, smallerNode *LinkedList

//     if countOne > countTwo {
//         biggerNode = currOne
//         smallerNode = currTwo
//     } else {
//         biggerNode = currTwo
//         smallerNode = currOne
//     }

//     for i := 0; i < diff; i++ {
//         biggerNode = biggerNode.Next
//     }

//     for biggerNode != smallerNode {
//         biggerNode = biggerNode.Next
//         smallerNode = smallerNode.Next
//     }

// 	return biggerNode
// }

// func abs (a int) int {
//     if a > 0 {
//         return a
//     }
//     return -a
// }
