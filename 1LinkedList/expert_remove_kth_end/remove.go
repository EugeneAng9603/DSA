// remove kth node from end
// 1,2,3,4,5,6,7,8,9 , k=4
// output is 1,2,3,4,5,7,8,9

package main

import (
	"fmt"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveKthNodeFromEnd(head *LinkedList, k int) {
	// Write your code here.

	// so second stops at 4
	first, second, counter := head, head, 1
	for counter <= k {
		second = second.Next
		counter += 1
	}
	if second != nil {
		fmt.Print(second.Value)
	}
	// means k = len of linked list = 10
	if second == nil {
		head.Value = head.Next.Value
		head.Next = head.Next.Next
		return
	}

	// so second stops at 9, first at 5
	for second.Next != nil {
		first = first.Next
		second = second.Next
	}
	// now first at 5, second at 9

	first.Next = first.Next.Next
	return
}
