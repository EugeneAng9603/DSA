package main

// import (
//     "fmt"
// )

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	// Write your code here.
	curr := linkedList
	for curr != nil {
		next := curr.Next
		for next != nil && curr.Value == next.Value {
			next = next.Next
		}
		curr.Next = next
		curr = next
	}
	return linkedList
}
