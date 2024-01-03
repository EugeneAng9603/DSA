// O(n), O(n)
// input := []int{1, 3, 3, 2, 4, 1, 5, 3, 2}
// 	expected := 9 bcuz 9x1

// In conlusion:
// 1. when top >= curr_height, means we found the right bound of top,
// so pop it, and next.top is left bound(bcuz it's the last smaller bound, can't go further left),
// calculate Area(larger than maxArea?)
// 1.1 THEN keep checking if next.top <= curr_height, do again for it. RMB they share
// the same right bound
// 2. when top < curr_height, do nth then add idx to stack then move on
// 3. once reach next of last height, height = 0, do the same pop and calculate until stack
// is empty

// Stack store left bound , when value < top, pop until left bound
// that's how we keep track of the smallest height left bound and
// larger stairs e.g. if [1,2,3,4,5,6,2], then pop6,5,4,3,2, so 1->2 = 7
package main

import (
	"fmt"
)

func LargestRectangleUnderSkyline(buildings []int) int {
	// Write your code here.
	stack := []int{}
	maxArea := 0

	extended := append(buildings, 0)
	for i := 0; i < len(extended); i++ {
		height := extended[i]
		// if small than top, pop the top, calculate maxArea
		for len(stack) != 0 && height <= buildings[stack[len(stack)-1]] {
			var stackIdx int
			stackIdx, stack = stack[len(stack)-1], stack[:len(stack)-1]
			pillarHeight := buildings[stackIdx]
			// if stack is empty, width is i
			width := i
			if len(stack) != 0 {
				width = i - stack[len(stack)-1] - 1
			}
			maxArea = max(width*pillarHeight, maxArea)

		}

		stack = append(stack, i)
		fmt.Print(i, stack, maxArea)
		fmt.Print("\n")
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
