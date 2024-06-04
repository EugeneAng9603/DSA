// find loop
// 3,4,5,6,7,8,9,4
// return 4

package main

import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func FindLoop(head *LinkedList) *LinkedList {
	// Write your code here.
	// cannot start at head, else it will start same
	slow, fast := head.Next, head.Next.Next
	for fast != slow {
		slow, fast = slow.Next, fast.Next.Next
	}

	slow = head

	for fast != slow {
		slow, fast = slow.Next, fast.Next
	}

	return slow
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
	list := addMany(&LinkedList{Value: 3}, []int{4, 5, 6, 7, 8, 9})
	// ll2 := addMany(&LinkedList{Value: 9}, []int{4, 5})
	// expected := addMany(&LinkedList{Value: 1}, []int{9, 2, 2})
	tail := list
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = list.Next
	actual := FindLoop(list)
	i := 0
	current := actual
	for i < 20 {
		fmt.Print(current.Value)
		current = current.Next
		i += 1
	}
}
