package main

func permute(nums []int) [][]int {
	var result [][]int
	heapPermute(len(nums), nums, &result)
	return result
}

// Heap's algorithm: efficient in-place permutation generator
func heapPermute(n int, nums []int, result *[][]int) {
	if n == 1 {
		// Make a copy of nums to avoid referencing the same slice
		perm := make([]int, len(nums))
		copy(perm, nums)
		*result = append(*result, perm)
		return
	}
	for i := 0; i < n; i++ {
		heapPermute(n-1, nums, result)
		if n%2 == 1 {
			// For odd n, swap first and last
			nums[0], nums[n-1] = nums[n-1], nums[0]
		} else {
			// For even n, swap i and last
			nums[i], nums[n-1] = nums[n-1], nums[i]
		}
	}
}
