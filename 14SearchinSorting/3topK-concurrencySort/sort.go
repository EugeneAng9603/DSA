package main

import (
	"container/heap"
	"fmt"
	"sync"
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

func topK(nums []int, k int) []int {
	h := &MinHeap{}
	heap.Init(h)

	for _, num := range nums {
		if h.Len() < k {
			heap.Push(h, num)
		} else if num > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, num)
		}
	}
	result := make([]int, h.Len())
	copy(result, *h)
	return result
}

func concurrentTopK(data []int, k int, chunkSize int) []int {
	numChunks := (len(data) + chunkSize - 1) / chunkSize
	var wg sync.WaitGroup
	results := make(chan []int, numChunks)

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunk := data[i:end]
		wg.Add(1)

		go func(chunk []int) {
			defer wg.Done()
			results <- topK(chunk, k)
		}(chunk)
	}

	wg.Wait()
	close(results)

	// Merge all chunk top-k results
	merged := []int{}
	for chunkTopK := range results {
		merged = append(merged, chunkTopK...)
	}
	return topK(merged, k)
}

func main() {
	arr := []int{1, 99, 5, 21, 50, 43, 3, 88, 13, 77, 6, 65, 33, 17, 100}
	k := 5
	chunkSize := 4
	result := concurrentTopK(arr, k, chunkSize)
	fmt.Printf("Top %d elements (concurrent): %v\n", k, result)
}
