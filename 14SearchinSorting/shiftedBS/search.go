// logn, 1

// same, just do the BS in the sorted half
// check sorted half by checking left > or < mid

// expected := 8
//
//	output := ShiftedBinarySearch([]int{45, 61, 71, 72, 73, 0, 1, 21, 33, 37}, 33)
//	require.Equal(t, expected, output)
package main

func ShiftedBinarySearch(array []int, target int) int {
	// Write your code here.
	return helper(array, target, 0, len(array)-1)
}

func helper(array []int, target, left, right int) int {
	for left <= right {
		mid := (left + right) / 2
		leftNum, rightNum := array[left], array[right]

		if target == array[mid] {
			return mid
			// means left to mid is sorted
		} else if leftNum <= array[mid] {
			// target too small and within mid and left, so look at left
			if target >= leftNum && target < array[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
			// means mid to right is sorted
		} else if leftNum > array[mid] {
			// check target within mid and right
			if target > array[mid] && target <= rightNum {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
