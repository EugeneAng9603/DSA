// O(t-s), O(t-s), which t is target, s is starting hand

// find probability that dealer will bust
// up to 3 decimal places
// target := 21
// 	startingHand := 15
// 	expected := 0.45

// 2-6 dealer standing, 7-10 dealer busting
// if 1, dealer keeps drawing cuz <=16 (dealer must reach target-4 or above which is 17), dealer draw
// so in this case answer = first draw bust + second draw bust + ...
// from second draw onwards,
// 0.4 + 0.1*0.5 = 0.45

// memo for those smaller than 17
// from + 1-10, if 17-21 not busted so 0, over 21 busted so 1,
// for smaller find memo or calculate recursion
package main

import (
	"math"
)

func BlackjackProbability(target int, startingHand int) float64 {
	// Write your code here.
	memo := make(map[int]float64)
	return math.Round(calculate(startingHand, target, memo)*1000) / 1000
}

// if bust, 1
// if stand, 0
// else, calculate total += 0.1 * prob
// cuz each card is 0.1 chance
func calculate(currHand int, target int, memo map[int]float64) float64 {
	if val, ok := memo[currHand]; ok {
		return val
	}
	if currHand > target {
		return 1
	}
	if currHand+4 >= target {
		return 0
	}

	totalProb := 0.0
	for i := 1; i < 11; i++ {
		totalProb += 0.1 * calculate(currHand+i, target, memo)
	}

	memo[currHand] = totalProb
	return totalProb
}
