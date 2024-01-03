// n, 1
package main

// import (
//
//	"fmt"
//
// )
// array := []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}
// actual := IsMonotonic(array)
// require.True(t, actual)

func IsMonotonic(array []int) bool {
	// Write your code here.
	isNonInc := true
	isNonDec := true

	for idx := 1; idx < len(array); idx++ {
		if array[idx] < array[idx-1] {
			isNonDec = false
		} else if array[idx] > array[idx-1] {
			isNonInc = false
		}
	}

	return isNonInc || isNonDec
}
