/*
*
Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.
Return the linked list sorted as well.
Example 1:
Input: 1->2->3->3->4->4->5
Output: 1->2->5
Example 2:
Input: 1->1->1->2->3
Output: 2->3

*
*/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// edge case
	if head == nil {
		return head
	}
	dummy := &ListNode{0, head}
	dummyptr := dummy
	curr := head
	for curr != nil {
		for curr.Next != nil && curr.Next.Val == curr.Val {
			curr = curr.Next
		}
		// means no duplicate
		if dummyptr.Next == curr {
			dummyptr = dummyptr.Next
		} else {
			// curr stop at [1,1,1,2,3,3,4] last 1
			dummyptr.Next = curr.Next
		}
		curr = curr.Next
		// fmt.Print(dummyptr.Val)
		// fmt.Print("\n")
	}
	return dummy.Next
}

// ---------------------------------------------------------------------------
// Use dummy and dup bool
// func deleteDuplicates(head *ListNode) *ListNode {
// 	dummy := &ListNode{-10, head}
// 	listptr := dummy
// 	curr := head
// 	for curr != nil {
// 		next := curr.Next
// 		dup := false
// 		for next != nil && next.Val == curr.Val {
// 			next = next.Next
// 			dup = true
// 		}
// 		if dup == false {
// 			listptr.Next = curr
// 			curr = curr.Next
// 		} else {
// 			listptr.Next = next
// 			curr = next.Next
// 		}
// 		listptr = listptr.Next
// 	}
// 	return dummy.Next
// }

func main() {
	list := &ListNode{Val: 1}
	list.Next = &ListNode{Val: 1}
	list.Next.Next = &ListNode{Val: 1}
	list.Next.Next.Next = &ListNode{Val: 2}
	list.Next.Next.Next.Next = &ListNode{Val: 3}
	list.Next.Next.Next.Next.Next = &ListNode{Val: 3}
	list.Next.Next.Next.Next.Next.Next = &ListNode{Val: 4}

	node := deleteDuplicates(list)
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
}

// --------------------------------------------------------------------------------------------------------------------------------
// create new linkedlist and add
//
//	func deleteDuplicates(head *ListNode) *ListNode {
//	    newList := &ListNode{Val:-10}
//	    listptr := newList
//	    curr := head
//	    for curr != nil {
//	        next := curr.Next
//	        count := 0
//	        for next != nil && next.Val == curr.Val{
//	            next = next.Next
//	            count += 1
//	        }
//	        if count == 0 {
//	            listptr.Next = curr
//	            curr = curr.Next
//	        } else {
//	            listptr.Next = next
//	            curr = next.Next
//	        }
//	        listptr = listptr.Next
//	    }
//	    return newList.Next
//	}
//
