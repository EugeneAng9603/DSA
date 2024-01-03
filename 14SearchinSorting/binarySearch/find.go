package main

import (
	"fmt"
)

func BinarySearch(array []int, target int) int {
	// Write your code here.
	start := 0
	end := len(array) - 1
	for start <= end {

		mid := (start + end) / 2

		fmt.Print(array[start], array[mid], array[end])
		fmt.Print("\n")

		if array[mid] == target {
			return mid
		} else if array[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}

	}
	return -1
}
