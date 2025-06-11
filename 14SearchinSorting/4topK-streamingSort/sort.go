package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

type TopKStream struct {
	k    int
	heap *MinHeap
}

func NewTopKStream(k int) *TopKStream {
	h := &MinHeap{}
	heap.Init(h)
	return &TopKStream{k: k, heap: h}
}

func (s *TopKStream) Add(num int) {
	if s.heap.Len() < s.k {
		heap.Push(s.heap, num)
	} else if num > (*s.heap)[0] {
		heap.Pop(s.heap)
		heap.Push(s.heap, num)
	}
}

func (s *TopKStream) GetTopK() []int {
	res := make([]int, len(*s.heap))
	copy(res, *s.heap)
	return res
}

func main() {
	stream := NewTopKStream(3)
	data := []int{10, 5, 20, 3, 50, 2, 11, 1}

	for _, num := range data {
		stream.Add(num)
		fmt.Printf("After %d, top 3: %v\n", num, stream.GetTopK())
	}
}
