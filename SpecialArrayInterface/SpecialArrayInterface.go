package main

import "fmt"

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {

	return helper(array, 1)
}

func helper(array SpecialArray, multiplier int) int {
	sum := 0
	for _, num := range array {
		if cast, ok := num.(SpecialArray); ok {
			sum += helper(cast, multiplier+1)
		} else if cast, ok := num.(int); ok {
			sum += cast
		}
	}
	return sum * multiplier
}

func main() {
	input := SpecialArray{
		5, 2,
		SpecialArray{7, -1},
		3,
		SpecialArray{
			6,
			SpecialArray{
				-13, 8,
			},
			4,
		},
	}
	output := ProductSum(input)

	fmt.Print(output)
}
