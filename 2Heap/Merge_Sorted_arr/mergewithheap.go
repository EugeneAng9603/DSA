package main

import (
	"container/heap"
	"fmt"
)

type entry struct {
	value int
	row   int
	col   int
}

type entryHeap []entry

func MergeSortedArrays(arrays [][]int) []int {
	pq := make(entryHeap, 0)

	for row, arr := range arrays {
		heap.Push(&pq, entry{arr[0], row, 0})
	}
	heap.Init(&pq)
	result := []int{}
	// fmt.Print(&pq)
	// pq start with {-124 2 0} {1 0 0} {-1 1 0} {3 3 0} , the first entry of each arr[0] in arrs
	for pq.Len() > 0 {
		e := heap.Pop(&pq).(entry)
		// fmt.Print(e) {-124 2 0} means val -124, row is 2, col is 0
		result = append(result, e.value)
		// if havent reach end of row, push new entry
		if e.col+1 < len(arrays[e.row]) {
			heap.Push(&pq, entry{arrays[e.row][e.col+1], e.row, e.col + 1})
		}
	}
	return result
}

func (pq entryHeap) Len() int {
	return len(pq)
}

func (pq entryHeap) Swap(a, b int) {
	pq[a], pq[b] = pq[b], pq[a]
}

func (pq entryHeap) Less(a, b int) bool {
	return pq[a].value < pq[b].value
}

func (pq *entryHeap) Push(x interface{}) {
	*pq = append(*pq, x.(entry))
}

// return type of any
func (pq *entryHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func main() {
	arrays := [][]int{
		{1, 5, 9, 21},
		{-1, 0},
		{-124, 81, 121},
		{3, 6, 12, 20, 150},
	}
	fmt.Print(MergeSortedArrays(arrays))
}
