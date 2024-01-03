package main

// var minHeap = NewMinHeap([]int{48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41})
// minHeap.Insert(76)
// require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))
// require.Equal(t, -5, minHeap.Peek())
// require.Equal(t, -5, minHeap.Remove())
// require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))
// require.Equal(t, 2, minHeap.Peek())
// require.Equal(t, 2, minHeap.Remove())
// require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))
// require.Equal(t, 6, minHeap.Peek())
// minHeap.Insert(87)
// require.Equal(t, true, isMinHeapPropertySatisfied(*minHeap))

type MinHeap []int

func NewMinHeap(array []int) *MinHeap {
	// Do not edit the lines below.
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap(array)
	return ptr
}

// O(n), O(1)
func (h *MinHeap) BuildHeap(array []int) {
	// Write your code here.
	// last parent node
	first := (len(array) - 2) / 2
	// siftdown check for all parent node starting from the last
	// bcuz we want to settle the subtree first
	// then siftdown for better algorithms
	// why first+1? bcuz we floor(first), might miss out
	// idx:0,1,2,3,4,5,6,7,8,9 , start from 4 to make sure dont miss out
	for currIdx := first + 1; currIdx >= 0; currIdx-- {
		h.siftDown(currIdx, len(array)-1)
	}
}

// O(nlogn) , O(1)
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

// O(logn), O(1)
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

// O(1), O(1)
func (h MinHeap) Peek() int {
	// Write your code here.
	if len(h) == 0 {
		return -1
	}
	return h[0]
}

// O(logn), O(1)
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

// O(logn), O(1)
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
