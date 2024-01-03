// Best: O(n), average, worst: O(n2)
// Space: O(1)
package main

// keep dividing into two subArr [...] [.......]
// for every index in array 2, insert it into array1 by swapping if smaller

func InsertionSort(array []int) []int {
	// Write your code here.
	// i to indicate the array1 last index(keep growing)
	for i := range array {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
	return array
}
