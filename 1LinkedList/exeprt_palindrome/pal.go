// 0,1,2,2,1,0 output: true

// Solution 1: Iterative, reverse second half and compare with first hald
// O(n), O(1)
package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func LinkedListPalindrome(head *LinkedList) bool {
	// Write your code here.
	// find mid using fast slow
	// fast will stop at mid.Next for odd or head of second half
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	reversedSecond := reverse(slow)
	firstHalf := head

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
