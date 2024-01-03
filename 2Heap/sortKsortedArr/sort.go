// nlogk, k

// No need remove any from array, because we are using heap (with k+1) values
// to pop min, insert new value, so not duplicate

// input := []int{3, 2, 1, 5, 4, 7, 6, 5}
// k := 3
// expected := []int{1, 2, 3, 4, 5, 5, 6, 7}

package main

func SortKSortedArray(array []int, k int) []int {
	// Write your code here.
	if len(array) == 0 || k == 0 {
		return array
	}
	heapArr := make([]int, min(k+1, len(array)))
	copy(heapArr, array[:min(k+1, len(array))])
	heap := NewMinHeap(heapArr)

	nextIdxToInsertElement := 0
	for i := k + 1; i < len(array); i++ {
		currMin := heap.Remove()
		array[nextIdxToInsertElement] = currMin
		nextIdxToInsertElement += 1

		currElement := array[i]
		heap.Insert(currElement)
	}

	// last step if only k or <k remaining
	for !heap.isEmpty() {
		currMin := heap.Remove()
		array[nextIdxToInsertElement] = currMin
		nextIdxToInsertElement += 1
	}

	return array
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// --------------------- Heap construction here -------------------//
type MinHeap []int

func NewMinHeap(array []int) *MinHeap {
	// Do not edit the lines below.
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap(array)
	return ptr
}

func (h *MinHeap) BuildHeap(array []int) {
	// Write your code here.
	// last parent node
	first := (len(array) - 2) / 2
	// siftdown check for all parent node starting from the last
	// why first+1? bcuz we floor(first), might miss out
	// idx:0,1,2,3,4,5,6,7,8,9 , start from 4 to make sure dont miss out
	for currIdx := first + 1; currIdx >= 0; currIdx-- {
		h.siftDown(currIdx, len(array)-1)
	}
}

func (h *MinHeap) siftDown(currIdx, endIdx int) {
	// Write your code here.
	childOneIdx := currIdx*2 + 1
	// if we reach bottom, we dont sift down
	for childOneIdx <= endIdx {
		// check if got second child
		childTwoIdx := -1
		if currIdx*2+2 <= endIdx {
			childTwoIdx = currIdx*2 + 2
		}
		// find smaller for potentia node to swap
		idxToSwap := childOneIdx
		if childTwoIdx != -1 && (*h)[childTwoIdx] < (*h)[childOneIdx] {
			idxToSwap = childTwoIdx
		}
		// now check potential node and currNode
		if (*h)[currIdx] > (*h)[idxToSwap] {
			h.swap(currIdx, idxToSwap)
			currIdx = idxToSwap
			childOneIdx = currIdx*2 + 1
		} else {
			return
		}
	}
}

func (h *MinHeap) siftUp() {
	// Write your code here.
	currIdx := h.length() - 1
	parentIdx := (currIdx - 1) / 2
	for currIdx > 0 {
		curr, parent := (*h)[currIdx], (*h)[parentIdx]
		if curr < parent {
			h.swap(currIdx, parentIdx)
			currIdx = parentIdx
			parentIdx = (currIdx - 1) / 2
		} else {
			return
		}
	}
}

func (h MinHeap) Peek() int {
	// Write your code here.
	if len(h) == 0 {
		return -1
	}
	return h[0]
}

func (h *MinHeap) Remove() int {
	// Write your code here.
	length := h.length()
	// index of first and last
	h.swap(0, length-1)
	toBeRemoved := (*h)[length-1]
	*h = (*h)[:length-1]
	h.siftDown(0, length-2)
	return toBeRemoved
}

func (h *MinHeap) Insert(value int) {
	// Write your code here.
	*h = append(*h, value)
	h.siftUp()
}

func (h *MinHeap) length() int {
	return len(*h)
}

func (h MinHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) isEmpty() bool {
	return len(*h) == 0
}
