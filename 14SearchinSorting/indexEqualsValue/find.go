// logn, 1
package main

// if the value < idx, eliminate left half
// cuz left half smaller than value and idx, cnt be answer
// To search for first index Equal values, check its prev
// if arr[prev] < prev, now must me the smallest valid
// if not, check left half

// input := []int{-5, -3, 0, 3, 4, 5, 9}
// expected := 3

func IndexEqualsValue(array []int) int {
	// Write your code here.
	left := 0
	right := len(array) - 1

	for left <= right {
		mid := (left + right) / 2
		midVal := array[mid]

		if midVal < mid {
			left = mid + 1
			// found and that's the first in array
		} else if midVal == mid && mid == 0 {
			return mid
			// if prev cmi, return now
		} else if midVal == mid && array[mid-1] < mid-1 {
			return mid
			// check left, cuz val > mid
		} else {
			right = mid - 1
		}
	}
	return -1
}
