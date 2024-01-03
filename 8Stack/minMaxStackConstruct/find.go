// ALL O(1), O(1)

package main

type MinMaxStack struct {
	// Write your code here.
	stack       []int
	minMaxStack []entry
}

type entry struct {
	min int
	max int
}

func NewMinMaxStack() *MinMaxStack {
	// Write your code here.
	return &MinMaxStack{}
}

// O(1), O(1)
func (stack *MinMaxStack) Peek() int {
	// Write your code here.
	return stack.stack[len(stack.stack)-1]
}

func (stack *MinMaxStack) Pop() int {
	// Write your code here.
	stack.minMaxStack = stack.minMaxStack[:len(stack.minMaxStack)-1]
	out := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	return out
}

func (stack *MinMaxStack) Push(number int) {
	// Write your code here.
	newMinMax := entry{min: number, max: number}
	if len(stack.minMaxStack) > 0 {
		lastMinMax := stack.minMaxStack[len(stack.minMaxStack)-1]
		newMinMax.min = min(lastMinMax.min, number)
		newMinMax.max = max(lastMinMax.max, number)
	}
	stack.minMaxStack = append(stack.minMaxStack, newMinMax)
	stack.stack = append(stack.stack, number)
}

func (stack *MinMaxStack) GetMin() int {
	// Write your code here.
	return stack.minMaxStack[len(stack.minMaxStack)-1].min
}

func (stack *MinMaxStack) GetMax() int {
	// Write your code here.
	return stack.minMaxStack[len(stack.minMaxStack)-1].max
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
