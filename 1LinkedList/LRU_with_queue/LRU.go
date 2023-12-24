// Head and Tail of Queue(doubly linked list) are empty node, easier to implement
// Head.Right is first element, Tail.Left is last element
// Queue struct is with length
package main

import (
	"fmt"
)

const SIZE = 5

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  HashMap
}

type HashMap map[string]*Node

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: HashMap{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Check(str string) {
	node := &Node{}
	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: str}
	}
	// Add to queue
	c.Add(node)
	// Add to Hash
	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("Remove: %s\n", n.Val)
	LeftNode := n.Left
	RightNode := n.Right

	RightNode.Left = LeftNode
	LeftNode.Right = RightNode
	c.Queue.Length -= 1
	delete(c.Hash, n.Val)
	return n
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("Add: %s\n", n.Val)

	temp := c.Queue.Head.Right
	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = temp
	temp.Left = n

	c.Queue.Length += 1

	// Check if exceede cache size
	if c.Queue.Length > SIZE {
		// temp := c.Queue.Tail.Left
		// c.Queue.Tail.Left = temp.Left
		// c.Queue.Tail.Right = c.Queue.Tail
		c.Remove(c.Queue.Tail.Left)
	}
}

// Keep showing the queue(doubly linked list) after each action
func (c *Cache) Display(str string) {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < q.Length-1 {
			fmt.Print("<-->")
		}
		node = node.Right
	}
	fmt.Print("]\n")
}

func main() {
	fmt.Println("START LRU CACHE.")
	cache := NewCache()
	for _, word := range []string{"a", "b", "c", "a", "d", "b", "c"} {
		cache.Check(word)
		cache.Display(word)
	}

}
