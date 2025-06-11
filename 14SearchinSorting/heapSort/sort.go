package main

import (
	"fmt"
)

func heapSort(array []int) []int {
	buildHeap(array)
	for endIdx := len(array) - 1; endIdx > 0; endIdx-- {
		swap(array, 0, endIdx)
		siftDown(0, endIdx-1, array)
	}
	return array
}

func buildHeap(array []int) {
	firstParentIdx := (len(array) - 2) / 2
	for currentIdx := firstParentIdx; currentIdx >= 0; currentIdx-- {
		siftDown(currentIdx, len(array)-1, array)
	}
}

func siftDown(currentIdx, endIdx int, heap []int) {
	childOneIdx := currentIdx*2 + 1
	for childOneIdx <= endIdx {
		childTwoIdx := -1
		if currentIdx*2+2 <= endIdx {
			childTwoIdx = currentIdx*2 + 2
		}

		var idxToSwap int
		if childTwoIdx != -1 && heap[childTwoIdx] > heap[childOneIdx] {
			idxToSwap = childTwoIdx
		} else {
			idxToSwap = childOneIdx
		}

		if heap[idxToSwap] > heap[currentIdx] {
			swap(heap, currentIdx, idxToSwap)
			currentIdx = idxToSwap
			childOneIdx = currentIdx*2 + 1
		} else {
			return
		}
	}
}

func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	arr := []int{3, 5, 1, 10, 2, 7}
	sorted := heapSort(arr)
	fmt.Println(sorted)
}
