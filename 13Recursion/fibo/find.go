// O(n), O(1)
// expected := 5
//
//	output := GetNthFib(6)
package main

func GetNthFib(n int) int {
	// Write your code here.
	prevTwo := []int{0, 1}
	counter := 3
	for counter <= n {
		next := prevTwo[0] + prevTwo[1]
		prevTwo = []int{prevTwo[1], next}
		counter += 1
	}

	if n > 1 {
		return prevTwo[1]
	}
	return prevTwo[0]
}
