// n,n

package main

// for every value get its left and right running product. for this question
// left is [1,5,5,20]
// right is [8,8,2,1]
// answer is left x right

// input := []int{5, 1, 4, 2}
// expected := []int{8, 40, 10, 20}

func ArrayOfProducts(array []int) []int {
	// Write your code here.
	left := make([]int, len(array))
	right := make([]int, len(array))
	result := make([]int, len(array))

	for i := range array {
		result[i] = 1
		left[i] = 1
		right[i] = 1
	}

	leftRunning := 1
	for i := 0; i < len(array); i++ {
		left[i] = leftRunning
		leftRunning *= array[i]
	}

	rightRunning := 1
	for i := len(array) - 1; i >= 0; i-- {
		right[i] = rightRunning
		rightRunning *= array[i]
	}

	for i := range array {
		result[i] = left[i] * right[i]
	}
	return result
}
