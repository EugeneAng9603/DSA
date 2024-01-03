// O(n), O(1)
// 1. Use the property of the 'pairing' idea
// 2. Keep track of the opening and closing from left to right,
// if open = closing, update max
// 3. If closing > openinig, reset to 0 for both.
// 4. Also need to come from right, because if opening
/// is found faster than closing ((())
// 5. So come from right, then if opening = closing, update max.
// openinig > closing reset

// "(()))("
// output = 4 --> (())
package main

func LongestBalancedSubstring(str string) int {
	// Write your code here.

	maxLength := 0
	// FROM LEFT FIRST
	openingCount := 0
	closingCount := 0

	for _, char := range str {
		if char == '(' {
			openingCount += 1
		} else {
			closingCount += 1
		}

		if openingCount == closingCount {
			maxLength = max(maxLength, openingCount*2)
			// means wrong case
		} else if closingCount > openingCount {
			openingCount = 0
			closingCount = 0
		}
	}

	// THEN FROM RIGHT
	openingCount = 0
	closingCount = 0

	for i := len(str) - 1; i >= 0; i-- {
		char := str[i]
		if char == '(' {
			openingCount += 1
		} else {
			closingCount += 1
		}

		if openingCount == closingCount {
			maxLength = max(maxLength, openingCount*2)
			// means wrong case
		} else if closingCount < openingCount {
			openingCount = 0
			closingCount = 0
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
