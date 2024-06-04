// Leetcode 121

// Example 1:

// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
// Example 2:

// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transactions are done and the max profit = 0.

package main

// -----------------------------------------------------------------------------------------------
// O(n), O(1)
// func maxProfit(prices []int) int {
//     runningMax := 0
//     result := 0
//     for i := len(prices) - 1; i >= 0; i-- {
//         if runningMax - prices[i] > result {
//             result = runningMax - prices[i]
//         }
//         if prices[i] > runningMax {
//             runningMax = prices[i]
//         }
//     }
//     return result
// }

// -----------------------------------------------------------------------------------------------

// O(n), O(1)
// 2 ptr method, left for buy, right for sell
// right to explore all nums, left to update to smallest possible num
func maxProfit(prices []int) int {
	left, right := 0, 1
	result := 0
	for right < len(prices) {
		if prices[left] < prices[right] {
			// update profit
			curr := prices[right] - prices[left]
			if curr > result {
				result = curr
			}
		} else {
			left = right
		}

		right += 1
	}
	return result
}
