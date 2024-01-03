// n, 1

// inverted from half
// 012345 become 210 543
// 0123456 become 210 3 654
// Use fast slow to find mid
// if odd, slow = mid; if even, slow = head of 2nd
// reverse for 1 and 2,
// connect 1 to 2 for even, connect 1 to mid to 2 for odd

package main

import (
	"fmt"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func InvertedBisection(head *LinkedList) *LinkedList {
	// Write your code here.
	// use fast and slow to split into two halves
	if head.Next == nil || head == nil {
		return head
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// if fast is nil, means even, else odd

	var newHead1, newHead2 *LinkedList
	// if odd
	if fast != nil {
		head2 := slow.Next
		head1 := head
		newHead2 = reverse(head2, nil)
		newHead1 = reverse(head1, slow)

		// connect newHead1 to newHead2
		curr := newHead1
		for curr.Next != nil {
			curr = curr.Next
		}
		// Link 1 to mid
		curr.Next = slow
		curr = slow
		curr.Next = newHead2
		// if even
	} else {
		head1 := head
		newHead1 = reverse(head1, slow)
		newHead2 = reverse(slow, nil)

		//
		curr := newHead1
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = newHead2
	}
	fmt.Print(newHead1, newHead2)
	return newHead1
}

func reverse(node, end *LinkedList) *LinkedList {
	curr := node
	var prev *LinkedList
	for curr != end {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// test := &LinkedList{Value: 0}
// 	test.Next = &LinkedList{Value: 1}
// 	test.Next.Next = &LinkedList{Value: 2}
// 	test.Next.Next.Next = &LinkedList{Value: 3}
// 	test.Next.Next.Next.Next = &LinkedList{Value: 4}
// 	test.Next.Next.Next.Next.Next = &LinkedList{Value: 5}

// 	output := InvertedBisection(test)

// 	actual := output.ToArray()
// 	expected := []int{2, 1, 0, 5, 4, 3}
// 	require.Equal(t, expected, actual)
