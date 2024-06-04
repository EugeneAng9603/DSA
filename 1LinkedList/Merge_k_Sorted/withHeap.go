package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(PQ, 0)
	for _, node := range lists {
		if node != nil {
			// pq = append(pq, node)
			// can use both
			heap.Push(&pq, node)
		}
	}
	// print pq , pq is 112
	for i := 0; i < len(pq); i++ {
		fmt.Print(pq[i].Val)
	}
	if len(pq) == 0 {
		return nil
	}
	heap.Init(&pq)

	headPtr := &ListNode{}
	dummyHead := headPtr

	for len(pq) > 0 {
		min := heap.Pop(&pq)
		fmt.Print("min is: ", min)
		fmt.Print("\n")
		// check if min is of interface type *ListNode
		// min is: &{1 0xc000014260}
		// minNode is : &{1 0xc000014260}
		// cast/convert min into the *ListNode type
		minNode := min.(*ListNode)
		fmt.Print("minNode is : ", minNode)
		fmt.Print("\n")
		headPtr.Next = minNode
		headPtr = headPtr.Next

		if minNode.Next != nil {
			heap.Push(&pq, minNode.Next)
		}
	}
	return dummyHead.Next
}

type PQ []*ListNode

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(a, b int) {
	pq[a], pq[b] = pq[b], pq[a]
}

func (pq PQ) Less(a, b int) bool {
	return pq[a].Val < pq[b].Val
}

func (pq *PQ) Push(nodeInterface interface{}) {
	node := nodeInterface.(*ListNode)
	*pq = append(*pq, node)
}

// return type of any
func (pq *PQ) Pop() interface{} {
	old := *pq
	lastNode := old[len(*pq)-1]
	*pq = old[:len(*pq)-1]
	return lastNode
}

func main() {
	lists := make([]*ListNode, 3)
	lists[0] = &ListNode{Val: 3}
	lists[0].Next = &ListNode{Val: 5}
	lists[0].Next.Next = &ListNode{Val: 8}

	lists[1] = &ListNode{Val: 1}
	lists[1].Next = &ListNode{Val: 4}
	lists[1].Next.Next = &ListNode{Val: 7}

	lists[2] = &ListNode{Val: 2}
	lists[2].Next = &ListNode{Val: 6}
	result := mergeKLists(lists)

	for result != nil {
		fmt.Print(result.Val)
		result = result.Next
	}
}
