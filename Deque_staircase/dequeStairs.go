package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

func StaircaseTraversal(height int, maxSteps int) int {
	var q deque.Deque[int]
	q.PushBack(1) // add value to dequeue, initiate with 1 for base case
	var windowSum = 0
	for i := 1; i < height+1; i++ {
		windowSum += q.Back() // back of dequeue
		if i > maxSteps {
			windowSum -= q.Front() //front of dequeue
			q.PopFront()
		}
		q.PushBack(windowSum)
	}

	return windowSum
}

func main() {
	fmt.Print(StaircaseTraversal(15, 5))
}
