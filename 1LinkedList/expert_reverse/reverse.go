package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func ReverseLinkedList(head *LinkedList) *LinkedList {
	// Write your code here.
	var prev, curr *LinkedList = nil, head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// --------------------------------------------------------------------
// Recursively get to the 5 to find newHead
// then going back, turn head.Next.Next to head, head to nil
// return the 5 as head
// once reach 5, return head
// back to 4, next(newHead) is 5
// so assign 4.Next.Next = 4 , now 5->4
// then 4->nil
// return 5->4->nil
// func ReverseLinkedList(head *LinkedList) *LinkedList {
// 	// Write your code here.
//     if head == nil || head.Next == nil {
//         return head
//     }
//     next := ReverseLinkedList(head.Next)
//     // link next to point itself
//     head.Next.Next = head
//     // itself point to nil
//     head.Next = nil
//     fmt.Print(next.Value)
//     fmt.Print("\n")

//     // return the new growing linkedlist
//     return next
// }
