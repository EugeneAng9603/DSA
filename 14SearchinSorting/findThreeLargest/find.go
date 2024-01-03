// O(n), O(1)
package main

// expected := []int{18, 141, 541}
// 	output := FindThreeLargestNumbers([]int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7,
import (
	"math"
)

func FindThreeLargestNumbers(array []int) []int {
	// Write your code here.
	threeNum := make([]int, 3)
	threeNum[0], threeNum[1], threeNum[2] = math.MinInt32, math.MinInt32, math.MinInt32

	for _, num := range array {
		if num > threeNum[2] {
			shiftUpdate(threeNum, num, 2)
		} else if num > threeNum[1] {
			shiftUpdate(threeNum, num, 1)
		} else if num > threeNum[0] {
			shiftUpdate(threeNum, num, 0)
		}
	}
	return threeNum
}

func shiftUpdate(arr []int, num, index int) {
	if index == 2 {
		arr[0], arr[1], arr[2] = arr[1], arr[2], num
	} else if index == 1 {
		arr[0], arr[1] = arr[1], num
	} else if index == 0 {
		arr[0] = num
	}
}
