package main

// O(n), O(n), NOT BEST
func MissingNumbers(nums []int) []int {
	// Write your code here.
	included := map[int]bool{}
	for _, num := range nums {
		included[num] = true
	}

	missing := make([]int, 0)
	for num := 1; num < len(nums)+3; num++ {
		if _, found := included[num]; !found {
			missing = append(missing, num)
		}
	}
	return missing
}

// ------better -------------------
// check algoexpert
