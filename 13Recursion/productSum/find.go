package main

// input := SpecialArray{
// 	5, 2,
// 	SpecialArray{7, -1},
// 	3,
// 	SpecialArray{
// 		6,
// 		SpecialArray{
// 			-13, 8,
// 		},
// 		4,
// 	},
// }
// output := ProductSum(input)
// expected := 12
// 5+2 + 2*(7-1) + 3 + 2*(6+ 3*(-13+8) +4)

// empty interfae, can be implemented by any type
// since it implements at least zero method
// in the input, it implements int type or specialArray type
type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	// Write your code here.
	return helper(array, 1)
}

func helper(array SpecialArray, multiplier int) int {
	sum := 0
	for _, num := range array {
		// .() is to check if num is an specialArray interface
		// where num is an interface{}
		if cast, ok := num.(SpecialArray); ok {
			// fmt.Print(num.(SpecialArray))
			// fmt.Print("\n")
			sum += helper(cast, multiplier+1)
			// check if num is int type
		} else if cast, ok := num.(int); ok {
			sum += cast
		}
	}
	return sum * multiplier
}
