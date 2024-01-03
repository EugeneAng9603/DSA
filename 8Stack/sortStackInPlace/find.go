// O(n2), O(n)
// sort stack without creating a new stack to store

// pop all by recursively calling function, then insert top recursively
package main

import "fmt"

func SortStack(stack []int) []int {
	// Write your code here.
	if len(stack) == 0 {
		return stack
	}

	top := stack[len(stack)-1]
	stack = stack[0 : len(stack)-1]
	SortStack(stack)

	InsertStack(&stack, top)
	return stack
}

// use pointer because adding to the same stack
func InsertStack(stack *[]int, value int) {
	if len(*stack) == 0 || (*stack)[len(*stack)-1] <= value {
		*stack = append(*stack, value)
		return
	}

	// if val is small, recursively pop the top until larger
	top := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	InsertStack(stack, value)
	// add the pop-ed top back
	*stack = append(*stack, top)
}

func main() {
	arr := []int{-5, 2, -2, 4, 3, 1}
	result := SortStack(arr)
	fmt.Print(result)
}
