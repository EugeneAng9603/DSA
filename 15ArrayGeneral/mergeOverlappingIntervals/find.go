// nlogn, n
package main

import (
	"sort"
)

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	// Write your code here.

	sorted := make([][]int, len(intervals))
	copy(sorted, intervals)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i][0] < sorted[j][0]
	})

	curr := sorted[0]
	merged := make([][]int, 0)
	merged = append(merged, curr)

	for _, next := range sorted {
		if curr[1] >= next[0] {
			curr[1] = max(curr[1], next[1])
		} else {
			curr = next
			merged = append(merged, curr)
		}
	}
	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// input := [][]int{
// 	{1, 2},
// 	{3, 5},
// 	{4, 7},
// 	{6, 8},
// 	{9, 10},
// }
// expected := [][]int{
// 	{1, 2},
// 	{3, 8},
// 	{9, 10},
// }
