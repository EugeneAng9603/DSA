package main

func FirstDuplicateValue(array []int) int {
	// Write your code here.
	for _, val := range array {
		val = abs(val)
		// -1 because from 1 to n
		if array[val-1] < 0 {
			return val
		}
		array[val-1] = array[val-1] * (-1)

	}
	return -1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
