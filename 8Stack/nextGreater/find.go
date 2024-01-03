// O(n), O(n)

// run arrayx2
// USE STACK
// add not-yet-found to stack by checking value < top e.g. [5,-3,-4]
// once found larger 6, pop all and upate result to 6 for all
package main

func NextGreaterElement(array []int) []int {
	// Write your code here.
	output := make([]int, 0)
	for range array {
		output = append(output, -1)
	}
	stack := make([]int, 0)

	for idx := 0; idx < 2*len(array); idx++ {
		circularIdx := idx % len(array)
		// found larger for all in stack, so pop all and update
		for len(stack) > 0 && array[circularIdx] > array[stack[len(stack)-1]] {
			var top int
			top, stack = stack[len(stack)-1], stack[0:len(stack)-1]
			output[top] = array[circularIdx]
		}
		stack = append(stack, circularIdx)
	}

	return output
}

// input := []int{2, 5, -3, -4, 6, 7, 2}
// 	expected := []int{5, 6, 6, 6, 7, -1, 5}
