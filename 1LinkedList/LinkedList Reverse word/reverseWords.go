package main

import (
	"fmt"
)

// This is an input struct. Do not edit.
type LinkedList struct {
	Value string
	Next  *LinkedList
}

func reverseWords(head *LinkedList) *LinkedList {

	dummy := &LinkedList{" ", head}
	pointer := dummy

	for pointer != nil {
		// skip " ", no need reverse
		for pointer.Next != nil && pointer.Next.Value == " " {
			pointer = pointer.Next
		}
		// now pointer starts at dummy, node is at "a" for "ab c def"
		node := pointer.Next
		counter := 0

		// find current token word length, e.g. ab then counter = 2
		for node != nil && node.Value != " " {
			node = node.Next
			counter++
		}
		// first letter of the token word
		curr := pointer.Next

		// if alr completed
		if curr == nil {
			break
		}

		// standard reverse here
		var next, prev *LinkedList
		for j := 0; j < counter; j++ {
			if curr == nil {
				break
			}
			next = curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}
		// then after reversing a group, handle the tail connection with pointer
		// prev at b->a , curr at " ", pointer still at dummy, now link dummy
		// tail of reversed token
		tail := pointer.Next // a
		tail.Next = curr     // a<-b , a.Next is " "
		pointer.Next = prev  //prev is at head of reversed token
		pointer = tail       // a , then carry on check for " c fed"
	}

	return dummy.Next
}

func addMany(linkedList *LinkedList, values []string) *LinkedList {
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

func getValues(linkedList *LinkedList) []string {
	values := []string{}
	current := linkedList
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	return values
}

// reverse k group
func main() {
	// ll1 := addMany(&LinkedList{Value: "a"}, []string{"b", "3", " ", "c", "5", " ", " ", "d", "e", "f", " "})
	ll1 := addMany(&LinkedList{Value: " "}, []string{"a", "b", " ", " ", "c", "!", "6", " ", "d", "e", "f", " "})
	actual := reverseWords(ll1)
	fmt.Print(getValues(actual))
}
