// 3,0,5,2,1,4 k=3 , for smaller than 3 go on the left, larger on the right
// output: 0,2,1,3,5,4

// keep track of the head, tail of smaller, equal, larger
package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RearrangeLinkedList(head *LinkedList, k int) *LinkedList {
	// Write your code here.
	var smallHead, smallTail, equalHead, equalTail, bigHead, bigTail *LinkedList

	curr := head
	for curr != nil {
		if curr.Value < k {
			smallHead, smallTail = grow(smallHead, smallTail, curr)
		} else if curr.Value == k {
			equalHead, equalTail = grow(equalHead, equalTail, curr)
		} else {
			bigHead, bigTail = grow(bigHead, bigTail, curr)
		}

		// Important step! break the link from prev to curr
		// so prev dont link back to curr
		prev := curr
		curr = curr.Next
		prev.Next = nil
	}

	// combine all 3
	// also need to handle if any node is nil, cannot nil.Next
	head1, tail1 := connect(smallHead, smallTail, equalHead, equalTail)
	finalHead, _ := connect(head1, tail1, bigHead, bigTail)
	//fmt.Print(smallHead)
	return finalHead
}

func connect(head1, tail1, head2, tail2 *LinkedList) (*LinkedList, *LinkedList) {
	newHead, newTail := head1, tail2
	// check final head
	if newHead == nil {
		newHead = head2
	}
	// check final tail
	if newTail == nil {
		newTail = tail1
	}
	// connect
	if tail1 != nil {
		tail1.Next = head2
	}
	return newHead, newTail
}

func grow(head, tail, node *LinkedList) (*LinkedList, *LinkedList) {
	newHead, newTail := head, node
	if newHead == nil {
		newHead = node
	}
	if tail != nil {
		tail.Next = node
	}
	return newHead, newTail
}
