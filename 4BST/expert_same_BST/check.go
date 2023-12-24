// given 2 array, check if both construct the same BST
// Do not construct to check
// arrayOne := []int{10, 15, 8, 12, 94, 81, 5, 2, 11}
// arrayTwo := []int{10, 8, 5, 15, 2, 12, 11, 94, 81}
// Output: true

// For every num in array, slice the following into 2 slices
// the smaller ones (become left sub) and the larger ones (right sub)
// the length and the first num must always be the same for same BST
// so no need deep equal the subarrays, recursively check length and first val

// EXTRA: if dont wanna keep passing array ( to save space)
// keep tracking of idx, so space: O(d) instead of n^2

package main

func SameBsts(arrayOne, arrayTwo []int) bool {
	// Write your code here.
	return helper(arrayOne, arrayTwo)
}

func helper(first, second []int) bool {
	if len(first) == 0 && len(second) == 0 {
		return true
	}

	if len(first) != len(second) || first[0] != second[0] {
		return false
	}

	smallerThanFirst := getSmaller(first)
	smallerThanSecond := getSmaller(second)
	largerThanFirst := getLarger(first)
	largerThanSecond := getLarger(second)

	return helper(smallerThanFirst, smallerThanSecond) &&
		helper(largerThanFirst, largerThanSecond)
}

func getSmaller(arr []int) []int {
	result := []int{}
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[0] {
			result = append(result, arr[i])
		}
	}
	return result
}

func getLarger(arr []int) []int {
	result := []int{}
	for i := 1; i < len(arr); i++ {
		if arr[i] >= arr[0] {
			result = append(result, arr[i])
		}
	}
	return result
}
