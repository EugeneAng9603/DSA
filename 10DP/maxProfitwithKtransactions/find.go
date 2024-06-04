// [5,11,3,50,60,90] , k=2
// buy 5, sell 11, buy 3, sell 90 makes most profit = 93

// rows is number of transactions 0,1,2
// cols is days 0,1,2,3,4,5,6
// initiate by all 0 for first row and first col

// for first row, max profit is max(currMax, max of currPrice - prev)
// for subsequent row, meaning >1 transaction
// we have 2 choice, either to sell it (if we bought before hand)
// or we dont sell, then max is max of prev day
// so profit[tran][day] = max of
// 1. prices[d] + max(-prices[x]+profit[t-1][x]) of all prev t-1 result
// 2. profit[t][d-1]
// reason is because if we sell today, means we will have bought on day x
// profit should be currPrice - prices[x] + profit[t-1][x]
// to find max profit, we must find max(-prices[x]+profit[t-1][x]) of all prev day
// so O(n2k)
// not good enough, so we can store the maxSoFar since we are repeating when x=0,1,2,3,4,5... for every check
// maxSoFar = max(maxSoFar, profit[t-1][d-1] - prices[d-1])
// now O(nk), O(nk)

// --------------------------------------------------------------------------------------------------------
// solution 2

// in fact we are only using prev row and maxSoFar, so just store 2 rows, then O(nk), O(n)
// just switch between first and second row

// code 2
// O(nk), O(n)
// only store 2 rows because we are only using prev and curr row
package main

import "math"

func MaxProfitWithKTransactions(prices []int, k int) int {
	// Write your code here.
	if len(prices) == 0 {
		return 0
	}
	evenProfit := make([]int, len(prices))
	oddProfit := make([]int, len(prices))
	var currProfit, prevProfit []int
	for t := 1; t < k+1; t++ {
		maxSoFar := math.MinInt32
		if t%2 == 1 {
			currProfit, prevProfit = oddProfit, evenProfit
		} else {
			currProfit, prevProfit = evenProfit, oddProfit
		}

		for d := 1; d < len(prices); d++ {
			maxSoFar = max(maxSoFar, prevProfit[d-1]-prices[d-1])
			currProfit[d] = max(currProfit[d-1], maxSoFar+prices[d])
		}
	}
	if k%2 == 0 {
		return evenProfit[len(prices)-1]
	}
	return oddProfit[len(prices)-1]
}

func max(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if num > curr {
			curr = num
		}
	}
	return curr
}

// Code 1
// O(nk), O(nk)
// package main

// import "math"

// func MaxProfitWithKTransactions(prices []int, k int) int {
// 	// Write your code here.
//     if len(prices) == 0 {
//         return 0
//     }

//     profit := make([][]int, k+1)
//     for i := range profit {
//         profit[i] = make([]int, len(prices))
//     }
//     for t := 1; t < k+1; t++ {
//         maxSoFar := math.MinInt32
//         for d := 1; d < len(prices); d++ {
//             maxSoFar = max(maxSoFar, profit[t-1][d-1] - prices[d-1])
//             profit[t][d] = max(profit[t][d-1], maxSoFar + prices[d])
//             // not sell, or sell
//             // if sell, means currPrice + profit[t-1][x] - price[x]
//             // which is maxSoFar = profit[t-1][x] - price[x]
//         }
//     }
//     return profit[k][len(prices)-1]

// }

// func max (arg int, rest ...int) int {
//     curr := arg
//     for _, num := range rest {
//         if num > curr {
//             curr = num
//         }
//     }
//     return curr
// }
