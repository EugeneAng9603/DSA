package main

import (
	"fmt"
)

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func reverseKgroup(head *LinkedList, k int) *LinkedList {
	node, cnt := head, 0
	for cnt < k {
		if node == nil {
			return head
		}
		node = node.Next
		cnt++
	}

	prev := reverseKgroup(node, k)
	for cnt > 0 {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
		cnt--
	}

	return prev
}

func addMany(linkedList *LinkedList, values []int) *LinkedList {
	current := linkedList
	for current.Next != nil {
		current = current.Next
	}
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

// reverse k group
func main() {
	ll1 := addMany(&LinkedList{Value: 2}, []int{3, 4, 5, 6, 7, 8})
	actual := reverseKgroup(ll1, 2)
	fmt.Print(getValues(actual))
}
