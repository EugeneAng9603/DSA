// n, n
package main

func SortedSquaredArray(array []int) []int {
	// Write your code here.
	result := make([]int, len(array))
	left := 0
	right := len(array) - 1
	for idx := len(array) - 1; idx >= 0; idx-- {
		if abs(array[left]) > abs(array[right]) {
			result[idx] = array[left] * array[left]
			left += 1
		} else {
			result[idx] = array[right] * array[right]
			right -= 1
		}
	}
	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
