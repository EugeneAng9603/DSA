// 0,1,2,2,1,0 output: true
// 1,2,3,3,2,1 slow and fast stop at (4,7)
// 1,2,3,4,3,2,1 slow and fast stop at (4,7)

// Solution 1: Iterative, reverse second half and compare with first hald
// O(n), O(1)
package main

import "fmt"

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func LinkedListPalindrome(head *LinkedList) bool {
	// Write your code here.
	// find mid using fast slow
	// slow will stop at mid for odd or head of second half for even
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	reversedSecond := reverse(slow)
	firstHalf := head
	// test1 := reversedSecond
	// test2 := firstHalf
	// for test1 != nil {
	// 	fmt.Print(test1.Value)
	// 	test1 = test1.Next
	// }
	// for test2 != nil {
	// 	fmt.Print(test2.Value)
	// 	test2 = test2.Next
	// }

	for reversedSecond != nil {
		if firstHalf.Value != reversedSecond.Value {
			return false
		} else {
			firstHalf = firstHalf.Next
			reversedSecond = reversedSecond.Next
		}
	}

	return true
}

func reverse(node *LinkedList) *LinkedList {
	var prev *LinkedList = nil
	curr := node
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
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
	ll := addMany(&LinkedList{Value: 1}, []int{2, 3, 4, 3, 2, 1})
	// ll2 := addMany(&LinkedList{Value: 9}, []int{4, 5})
	// expected := addMany(&LinkedList{Value: 1}, []int{9, 2, 2})
	actual := LinkedListPalindrome(ll)
	fmt.Print(actual)
}
