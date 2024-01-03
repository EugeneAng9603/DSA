// Best: O(n), average and worst: O(n2)
// space: O(1)
package main

func BubbleSort(array []int) []int {
	// Write your code here.
	counter := 0
	isSorted := false
	// use isSorted to indicate if there's swap, if no swap at all
	// means already done

	for !isSorted {
		isSorted = true
		for i := 0; i < len(array)-1-counter; i++ {
			if array[i] > array[i+1] {
				swap(array, i, i+1)
				isSorted = false
			}
		}
		// use counter to reduce the time of check needed
		counter += 1
	}

	return array
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
