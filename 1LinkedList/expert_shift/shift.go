// 0,1,2,3,4,5 k=2
// 4,5,0,1,2,3

// find tail position first
package main

import (
	"fmt"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func ShiftLinkedList(head *LinkedList, k int) *LinkedList {
	// Write your code here.
	length := 1
	listTail := head
	for listTail.Next != nil {
		length += 1
		listTail = listTail.Next
	}

	offset := abs(k) % length
	fmt.Print(length, offset)
	// no effort, return head
	if offset == 0 {
		return head
	}

	// if its 2, tail should be at 3 (position 4)
	// if its -2, tail be at 1 (position 2)
	newTailPosition := length - offset
	if k < 0 {
		newTailPosition = offset
	}

	// reach newTail
	newTail := head
	for i := 1; i < newTailPosition; i++ {
		newTail = newTail.Next
	}

	// declare new Head
	newHead := newTail.Next
	// tail -> nil
	newTail.Next = nil
	// when find length of ll
	listTail.Next = head
	return newHead
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
