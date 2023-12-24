// Feel free to add methods and fields to the struct definitions.

package main

type Node struct {
	Value      int
	Prev, Next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	// Write your code here.
	return &DoublyLinkedList{}
}

func (ll *DoublyLinkedList) SetHead(node *Node) {
	// Write your code here.
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}
	ll.InsertBefore(ll.Head, node)
}

func (ll *DoublyLinkedList) SetTail(node *Node) {
	// Write your code here.
	if ll.Tail == nil {
		ll.SetHead(node)
		return
	}
	ll.InsertAfter(ll.Tail, node)
}

func (ll *DoublyLinkedList) InsertBefore(node, nodeToInsert *Node) {
	// Write your code here.
	// Edge: if already in l1 and is head of tail, return
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return
	}

	// else remove the node in ll
	ll.Remove(nodeToInsert)
	// newNode point to left right first
	nodeToInsert.Prev = node.Prev
	nodeToInsert.Next = node
	// check nil, cuz cant nil.Next
	if node.Prev == nil {
		ll.Head = nodeToInsert
	} else {
		node.Prev.Next = nodeToInsert
	}
	node.Prev = nodeToInsert
}

func (ll *DoublyLinkedList) InsertAfter(node, nodeToInsert *Node) {
	// Write your code here.
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return
	}

	ll.Remove(nodeToInsert)
	nodeToInsert.Prev = node
	nodeToInsert.Next = node.Next
	if node.Next == nil {
		ll.Tail = nodeToInsert
	} else {
		node.Next.Prev = nodeToInsert
	}
	node.Next = nodeToInsert
}

func (ll *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *Node) {
	// Write your code here.
	if position == 1 {
		ll.SetHead(nodeToInsert)
		return
	}
	node := ll.Head
	currPosition := 1
	for node != nil && currPosition != position {
		node = node.Next
		currPosition += 1
	}
	if node != nil {
		ll.InsertBefore(node, nodeToInsert)
	} else {
		ll.SetTail(nodeToInsert)
	}
}

func (ll *DoublyLinkedList) RemoveNodesWithValue(value int) {
	// Write your code here.
	node := ll.Head
	for node != nil {
		nodeToRemove := node
		node = node.Next
		if nodeToRemove.Value == value {
			ll.Remove(nodeToRemove)
		}
	}
}

func (ll *DoublyLinkedList) Remove(node *Node) {
	// Write your code here.
	if node == ll.Head {
		ll.Head = ll.Head.Next
	}
	if node == ll.Tail {
		ll.Tail = ll.Tail.Prev
	}
	ll.RemoveBindings(node)
}

func (ll *DoublyLinkedList) ContainsNodeWithValue(value int) bool {
	// Write your code here.
	node := ll.Head
	if node != nil && node.Value != value {
		node = node.Next
	}
	return node != nil
}

func (ll *DoublyLinkedList) RemoveBindings(node *Node) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	node.Prev = nil
	node.Next = nil
}

// ------------------------------------------------------------------------------------------------------
// This file is initialized with a code version of this
// question's sample test case. Feel free to add, edit,
// or remove test cases in this file as you see fit!

// package main

// import (
// 	"github.com/stretchr/testify/require"
// )

// func (s *TestSuite) TestCase1(t *TestCase) {
// 	linkedList := NewDoublyLinkedList()
// 	one := NewNode(1)
// 	two := NewNode(2)
// 	three := NewNode(3)
// 	three2 := NewNode(3)
// 	three3 := NewNode(3)
// 	four := NewNode(4)
// 	five := NewNode(5)
// 	six := NewNode(6)
// 	bindNodes(one, two)
// 	bindNodes(two, three)
// 	bindNodes(three, four)
// 	bindNodes(four, five)
// 	linkedList.Head = one
// 	linkedList.Tail = five

// 	linkedList.SetHead(four)
// 	require.Equal(t, getNodeValuesHeadToTail(linkedList), []int{4, 1, 2, 3, 5})
// 	require.Equal(t, getNodeValuesTailToHead(linkedList), []int{5, 3, 2, 1, 4})

// 	linkedList.SetTail(six)
// 	require.Equal(t, getNodeValuesHeadToTail(linkedList), []int{4, 1, 2, 3, 5, 6})
// 	require.Equal(t, getNodeValuesTailToHead(linkedList), []int{6, 5, 3, 2, 1, 4})

// 	linkedList.InsertBefore(six, three)
// 	require.Equal(t, getNodeValuesHeadToTail(linkedList), []int{4, 1, 2, 5, 3, 6})
// 	require.Equal(t, getNodeValuesTailToHead(linkedList), []int{6, 3, 5, 2, 1, 4})

// 	linkedList.InsertAfter(six, three2)
// 	require.Equal(t, getNodeValuesHeadToTail(linkedList), []int{4, 1, 2, 5, 3, 6, 3})
// 	require.Equal(t, getNodeValuesTailToHead(linkedList), []int{3, 6, 3, 5, 2, 1, 4})

// 	linkedList.InsertAtPosition(1, three3)
// 	require.Equal(t, getNodeValuesHeadToTail(linkedList), []int{3, 4, 1, 2, 5, 3, 6, 3})
// 	require.Equal(t, getNodeValuesTailToHead(linkedList), []int{3, 6, 3, 5, 2, 1, 4, 3})

// 	linkedList.RemoveNodesWithValue(3)
// 	require.Equal(t, getNodeValuesHeadToTail(linkedList), []int{4, 1, 2, 5, 6})
// 	require.Equal(t, getNodeValuesTailToHead(linkedList), []int{6, 5, 2, 1, 4})

// 	linkedList.Remove(two)
// 	require.Equal(t, getNodeValuesHeadToTail(linkedList), []int{4, 1, 5, 6})
// 	require.Equal(t, getNodeValuesTailToHead(linkedList), []int{6, 5, 1, 4})

// 	require.Equal(t, linkedList.ContainsNodeWithValue(5), true)
// }

// func NewNode(value int) *Node { return &Node{Value: value} }

// func getNodeValuesHeadToTail(ll *DoublyLinkedList) []int {
// 	values := []int{}
// 	node := ll.Head
// 	for node != nil {
// 		values = append(values, node.Value)
// 		node = node.Next
// 	}
// 	return values
// }

// func getNodeValuesTailToHead(ll *DoublyLinkedList) []int {
// 	values := []int{}
// 	node := ll.Tail
// 	for node != nil {
// 		values = append(values, node.Value)
// 		node = node.Prev
// 	}
// 	return values
// }

// func bindNodes(nodeOne *Node, nodeTwo *Node) {
// 	nodeOne.Next = nodeTwo
// 	nodeTwo.Prev = nodeOne
// }
