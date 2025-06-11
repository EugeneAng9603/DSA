// Question: https://leetcode.com/problems/find-median-from-data-stream/description/

// 🔍 Time & Space Complexity
// Operation	Time	Space
// AddNum()	O(log n)	O(n)
// FindMedian()	O(1)	O(1)

package main

import (
	"container/heap"
	"fmt"
)

// MinHeap (for the larger half)
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MaxHeap (for the smaller half)
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // reverse for max-heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MedianFinder struct
type MedianFinder struct {
	low  *MaxHeap
	high *MinHeap
}

func Constructor() MedianFinder {
	low := &MaxHeap{}
	high := &MinHeap{}
	heap.Init(low)
	heap.Init(high)
	return MedianFinder{low: low, high: high}
}

func (m *MedianFinder) AddNum(num int) {
	heap.Push(m.low, num)
	// maintain the invariant: max(low) <= min(high)
	heap.Push(m.high, heap.Pop(m.low))
	if m.high.Len() > m.low.Len() {
		heap.Push(m.low, heap.Pop(m.high))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.low.Len() > m.high.Len() {
		return float64((*m.low)[0])
	}
	return (float64((*m.low)[0]) + float64((*m.high)[0])) / 2.0
}

func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	fmt.Println(obj.FindMedian()) // 1.5
	obj.AddNum(3)
	fmt.Println(obj.FindMedian()) // 2.0
	obj.AddNum(4)
	fmt.Println(obj.FindMedian()) // 2.5
	obj.AddNum(6)
	fmt.Println(obj.FindMedian()) // 3.0
	obj.AddNum(0)
	fmt.Println(obj.FindMedian()) // 2.5
}
