package main

import (
	"fmt"
)

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func SumOfLinkedLists(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
	// Write your code here.
	newHead := &LinkedList{Value: 0}
	curr := newHead

	llOne := linkedListOne
	llTwo := linkedListTwo
	carry := 0
	for llOne != nil || llTwo != nil || carry != 0 {
		var valueOne, valueTwo int
		if llOne != nil {
			valueOne = llOne.Value
		}
		if llTwo != nil {
			valueTwo = llTwo.Value
		}
		sum := valueOne + valueTwo + carry

		newValue := sum % 10
		newNode := &LinkedList{Value: newValue}
		curr.Next = newNode
		curr = newNode

		//carry
		carry = sum / 10

		// check if need to move to next
		if llOne != nil {
			llOne = llOne.Next
		}
		if llTwo != nil {
			llTwo = llTwo.Next
		}
	}
	return newHead.Next
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
	ll1 := addMany(&LinkedList{Value: 2}, []int{4, 7, 1})
	ll2 := addMany(&LinkedList{Value: 9}, []int{4, 5})
	// expected := addMany(&LinkedList{Value: 1}, []int{9, 2, 2})
	actual := SumOfLinkedLists(ll1, ll2)
	fmt.Print(getValues(actual))
}
