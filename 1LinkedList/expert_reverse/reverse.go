package main

import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// func ReverseLinkedList(head *LinkedList) *LinkedList {
// 	// Write your code here.
// 	var prev, curr *LinkedList = nil, head
// 	for curr != nil {
// 		next := curr.Next
// 		curr.Next = prev
// 		prev = curr
// 		curr = next
// 	}
// 	return prev
// }

// --------------------------------------------------------------------
// Recursively get to the 5 to find newHead
// once reached, turn 5->4 , 4->nil
// return the same newHead
func ReverseLinkedList(head *LinkedList) *LinkedList {
	// Write your code here.
	if head == nil || head.Next == nil {
		return head
	}
	// first bring 2 to reverse, should get back 5>4>3>2, while 1>2)
	next := ReverseLinkedList(head.Next)
	// link next to point itself (so 5->4)
	head.Next.Next = head
	// itself point to nil (so 5->4>nil, and 1->2->3->4)
	head.Next = nil
	// fmt.Print(next.Value)
	// fmt.Print("\n")

	// return the head of the new growing linkedlist
	return next
}

func addMany(linkedList *LinkedList, values []int) *LinkedList {
	current := linkedList
	// for current.Next != nil {
	// 	current = current.Next
	// }
	for _, value := range values {
		current.Next = &LinkedList{Value: value}
		current = current.Next
	}
	return linkedList
}

func getValues(linkedList *LinkedList) []int {
	values := []int{}
	current := linkedList
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	return values
}

func main() {
	ll1 := addMany(&LinkedList{Value: 1}, []int{2, 3, 4, 5, 6, 7})
	// expected := addMany(&LinkedList{Value: 1}, []int{9, 2, 2})
	actual := ReverseLinkedList(ll1)
	fmt.Print(getValues(actual))
}
