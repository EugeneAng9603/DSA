// O(n), O(n)
// "462839" , numDigits = 2
// output is "6839", cuz removing 4 and 2 can give largest output

package main

// Add all number into stack one by one, if newNum is larger, pop out
// prev until numDigits = 0 or until newNum not larger
// if already largest, simply pop numDigits number of times
// so always led by the larger numbers
func BestDigits(number string, numDigits int) string {
	// Write your code here.
	stack := make([]rune, 0)

	for _, digit := range number {
		// 6 > 4, pop
		for numDigits > 0 && len(stack) > 0 && digit > stack[len(stack)-1] {
			stack = stack[0 : len(stack)-1]
			numDigits -= 1
		}
		// append every new digit
		stack = append(stack, digit)
	}

	// if longer than answer, then just pop the remaining
	for numDigits > 0 {
		numDigits -= 1
		stack = stack[0 : len(stack)-1]
	}

	return string(stack)
}
