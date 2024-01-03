// dp method, however O(n2), O(n)
// expected := 4
// 	output := MinNumberOfJumps([]int{3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3})
// 3 --> (4 or 2) -> (2 or 3) -> 7 -> 3

// in the array, we are storing the min num of jumps to go from 0 to itself
// so for every i, check j from 0 to i
// if can jump, array[j] + j >= i
// then jumps[i] = min(jumps[i], jumps[j] + 1) means 0toj,then +1 jump to i

package main

func MinNumberOfJumps(array []int) int {
	// Write your code here.
	// min jumps from 0 to current position
	dp := make([]int, len(array))
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	for i := 1; i < len(array); i++ {
		for j := 0; j < i; j++ {
			// if left + idx > i, means can reach
			// AND if never updated dp[i] or dp[j] + 1 is new min
			// then update
			if array[j] >= i-j && (dp[i] == -1 || dp[j]+1 < dp[i]) {
				dp[i] = dp[j] + 1
			}
		}
	}

	return dp[len(array)-1]
}

// better -----------------------------------------------------
// O(n), O(1)
// use some ideology to help us
// if we know the maxReach e.g. index 1, value 4, maxReach is 5
// and we know the steps array[i]
// every iteration we remaining step -= 1
// if steps == 0, means must jump, jumpCount += 1
// then steps = maxReach - i for current steps if steps == 0
// (this means the steps remaining from i)

// package main

// import(
//     "fmt"
// )

// func MinNumberOfJumps(array []int) int {
// 	// Write your code here.
//     if len(array) == 1 {
//         return 0
//     }

//     jumps := 0
//     maxReach := array[0]
//     steps := array[0]

//     // no need jump form last index, so len(arr) - 1
//     for i := 1; i < len(array) - 1; i++ {
//         maxReach = max(maxReach, array[i] + i)
//         steps -= 1 //consume a step

//         // check if step reach 0, to make a new jump
//         if steps == 0 {
//             jumps += 1
//             steps = maxReach - i
//         }
//         fmt.Print(maxReach, steps, jumps)
//         fmt.Print("\n")
//     }
// 	return jumps + 1
// }

// func max(a, b int) int{
//     if a > b {
//         return a
//     }
//     return b
// }
