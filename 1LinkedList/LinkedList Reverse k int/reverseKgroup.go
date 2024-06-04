package main

import (
	"fmt"
)

// 1->2->3->4->5->6->7->8
// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

// Recursive
func reverseKgroup(head *LinkedList, k int) *LinkedList {
	node, cnt := head, 0
	for cnt < k {
		if node == nil {
			return head
		}
		node = node.Next
		cnt++
	}

	prev := reverseKgroup(node, k) // prev is at 7, head is 4
	for cnt > 0 {
		//
		next := head.Next // next at 5
		head.Next = prev  // 4->7
		prev = head       // prev at 4
		head = next       // head at 5, hence 5->6->7 , 4->7
		cnt--
	}

	return prev
}

// ------------------------------------------------------------------------------------------
// Iterative
/* Iterative Solution */

// func reverseKGroup(head *ListNode, k int) *ListNode {
//     if head == nil || k == 1 {
//         return head
//     }

//     l := length(head)
//     preHead := &ListNode{Next: head}

//     prev := preHead
//     for l >= k {
//         curr := prev.Next
//         for i := 1; i < k; i++ {
//             next := curr.Next
//             curr.Next = next.Next
//             next.Next = prev.Next
//             prev.Next = next
//         }
//         prev = curr
//         l -= k
//     }

//     return preHead.Next
// }

// func length(head *ListNode) int {
//     count := 0

//     for head != nil {
//         head = head.Next
//         count++
//     }
//     return count
// }

func addMany(linkedList *LinkedList, values []int) *LinkedList {
	current := linkedList
	for current.Next != nil {
		current = current.Next
	}
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

// reverse k group
func main() {
	ll1 := addMany(&LinkedList{Value: 2}, []int{3, 4, 5, 6, 7, 8})
	actual := reverseKgroup(ll1, 2)
	fmt.Print(getValues(actual))
}
