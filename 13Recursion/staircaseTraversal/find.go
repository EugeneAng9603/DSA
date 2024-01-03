// height=4 , max=2, output = 5
// 10, 2, output = 89
// Solution 1: just do recursion
// O(k^n), bad

package main

func StaircaseTraversal(height int, maxSteps int) int {
	// Write your code here.
	return helper(height, maxSteps)
}

// for maxstep = 3
// helper(6) = height(5) + height(4) + height(3)
// for height(5) same logic.
func helper(height, maxSteps int) int {
	if height <= 1 {
		return 1
	}

	var ways = 0

	// height might drop to less than step
	for step := 1; step < min(maxSteps, height)+1; step++ {
		ways += helper(height-step, maxSteps)
	}
	return ways
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// memo -------------------------------------
// Solution 2: Also recursion, but with memo
// O(kn), O(n)

// package main

// func StaircaseTraversal(height int, maxSteps int) int {
// 	// Write your code here.
//     table := make(map[int]int, 0)
//     table[0] = 1
//     table[1] = 1

// 	return helper(height, maxSteps, table)
// }

// func helper(height, maxSteps int, memo map[int]int) int {
//     if ways, found := memo[height]; found {
//         return ways
//     }

//     var numberOfWays = 0
//     for step := 1; step < min(maxSteps, height) + 1; step++{
//         numberOfWays += helper(height - step, maxSteps, memo)
//     }
//     memo[height] = numberOfWays
//     return numberOfWays
// }

// func min (a, b int) int {
//     if a > b {
//         return b
//     }
//     return a
// }

// ------------------------------------------------------------
// Solution 4: use dequeue , no need track ways[n] array
// O(n), O(k) maxSteps size

// package main

// import (
//     "fmt"
//     "github.com/gammazero/deque"
// )

// func StaircaseTraversal(height int, maxSteps int) int {
//     var q deque.Deque[int]
//     q.PushBack(1)
//     var windowSum = 0
//     for i := 1; i < height + 1; i++ {
//         windowSum += q.back()
//         if i > maxSteps {
//             windowSum -= window.front()
//             window.PopFront()
//         }
//         window.PushBack(windowSum)
//     }

//     return windowSum
// }

// option 3 ------------------------------------------------------------
// Solution 3: Sliding window (2 index)
// O(n), O(n)

// package main

// import (
//     "fmt"
// )

// e.g. 10 = 9 + 8 + 7
// start with 1, then 1 to 2, then 1 to 3
// next remove 1 and add 4 becomes 2 to 4
// reached 7 to 9, done
// [-5 0][-4 1][-3 2][-2 3][-1 4]
// [0 5][1 6][2 7][3 8][4 9][5 10]
// [6 11][7 12][8 13][9 14]

// func StaircaseTraversal(height int, maxSteps int) int {
// 	// Write your code here.
//     currWays := 0
//     ways := []int{1}

//     for currHeight := 1; currHeight < height + 1; currHeight++ {
//         start := currHeight - maxSteps - 1
//         end := currHeight - 1
//         // start and end always of "maxSteps" diff
//         fmt.Print([]int{start, end})
//         // Keep adding the newest and removing the oldest
//         // If less than 3 items, then just add without removing any
//         if start >= 0 {
//             currWays -= ways[start]
//         }

//         // always add the newest ways[n]
//         currWays += ways[end]
//         ways = append(ways, currWays)
//     }

// 	return ways[height]
// }
