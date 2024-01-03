// O(n), O(1)
// check if start from any index, can jump back to itself
// input := []int{2, 3, 1, -4, -4, 2}
// 	output := HasSingleCycle(input)
// 	expected := true

package main

func HasSingleCycle(array []int) bool {
	// Write your code here.
	visited := 0
	currIdx := 0
	for visited < len(array) {
		if visited > 0 && currIdx == 0 {
			return false
		}
		currIdx = getNext(currIdx, array)
		visited += 1
	}

	return currIdx == 0
}

func getNext(idx int, arr []int) int {
	next := (idx + arr[idx]) % len(arr)
	if next < 0 {
		return next + len(arr)
	} else {
		return next
	}
}
