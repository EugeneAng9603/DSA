package main

import (
	"fmt"
)

// array := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}
// actual := LongestPeak(array)
// require.Equal(t, 6, actual)

func LongestPeak(array []int) int {
	// Write your code here.
	longest := 0
	i := 1
	for i < len(array)-1 {
		isPeak := array[i] > array[i-1] && array[i] > array[i+1]
		if !isPeak {
			i++
			continue
		}

		left := i - 2
		for left >= 0 && array[left] < array[left+1] {
			left -= 1
		}

		right := i + 2
		for right <= len(array)-1 && array[right] < array[right-1] {
			right += 1
		}
		fmt.Print(left, right)
		// bcuz at this point, right will be at [0,1,2,3,1,0] 6th,
		// left will be at -1th, so right - left - 1 = 6
		longest = max(longest, right-left-1)
		// curr := right - left - 1
		// if curr > longest {
		//     longest = curr
		// }
		i = right
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
