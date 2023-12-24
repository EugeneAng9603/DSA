// 2,7,3,5 output: 3,5
package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func MiddleNode(linkedList *LinkedList) *LinkedList {
	// Write your code here.
	fast := linkedList
	slow := linkedList

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
