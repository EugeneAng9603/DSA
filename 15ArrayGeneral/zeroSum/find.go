// n,n
// input := []int{4, 2, -1, -1, 3}
// 	expected := true
// can find, 2, -1, -1 is 0

package main

func ZeroSumSubarray(nums []int) bool {
	// Write your code here.

	// Track each sums from [0,i]
	// Then if repeated, return true
	// e.g. [0,2] = 1, [0,5] = 1, means [3,5] = 0
	sums := map[int]bool{0: true}
	currSum := 0
	for _, num := range nums {
		currSum += num
		if _, sumIsInSet := sums[currSum]; sumIsInSet {
			return true
		}
		sums[currSum] = true
	}
	return false
}
