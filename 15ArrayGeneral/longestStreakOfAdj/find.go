// n,1

// find the index of 0, which if replaced by 1, can form the longest streak of 1
// for every 1, just add on to get longest curr streak
// when reach 0, try update index and currLongest
// so we can just do one pass from left to right
// array := []int{1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1}
// actual := LongestStreakOfAdjacentOnes(array)
// require.Equal(t, 8, actual)
// }
package main

func LongestStreakOfAdjacentOnes(array []int) int {
	// Write your code here.
	tryZeroIdx := -1
	longest := 0
	currStreak := 0
	answer := -1
	for i := 0; i < len(array); i++ {
		if array[i] == 1 {
			currStreak += 1
		} else {
			// means from the left, the longest is as updated
			currStreak = i - tryZeroIdx
			tryZeroIdx = i
		}
		// check longest
		if currStreak > longest {
			longest = currStreak
			answer = tryZeroIdx
		}
	}
	return answer
}
