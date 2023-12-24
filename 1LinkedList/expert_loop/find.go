// find loop
// 3,4,5,6,7,8,9,4
// return 4

package main

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
