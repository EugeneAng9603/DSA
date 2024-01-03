package main

//import "fmt"

func TwoNumberSum(array []int, target int) []int {
	// Write your code here.

	// nums := map[int]bool{}
	nums := make(map[int]bool)
	for _, num := range array {
		potentialMatch := target - num
		if _, found := nums[potentialMatch]; found {
			// fmt.Println([]int{potentialMatch, num})
			return []int{potentialMatch, num}
		}
		nums[num] = true
	}
	return []int{}
}
