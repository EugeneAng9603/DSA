// Have the function MathChallenge(str) take the str parameter,
// which will be a simple mathematical formula with three numbers,
// a single operator (+, -, *, or /) and an equal sign (=) and
// return the two digits that complete the equation.

// In two of the numbers in the equation, there will be a single ? character,
// and your program should determine what digits are missing and return them
// separated by a space. For example, if str is "38?5 * 3 = 1?595" then your program should output 6 1.

// The ? character will always appear in both the first number and the last number
// in the mathematical expression. There will always be a unique solution.
// Examples
// Input: "56? * 106 = 5?678"
// Output: 3 9
// Input: "18?1 + 9 = 189?"
// Output: 8 0
// Input: "50?050 / 5 = ?01010"
// Output: 5 1

// Explanation:
// replaceQuestionMark: This helper function replaces the '?' in a string with a given digit.
// evaluateExpression: This function evaluates the expression with the given operator to check if it holds true.
// findMissingDigits: This function iterates through all possible digit combinations for
// the two numbers with '?' and checks if replacing the '?' with these digits makes the equation valid.
// MathChallenge: This function parses the input string into components and calls findMissingDigits to get the missing digits.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func replaceQuestionMark(s string, digit int) string {
	return strings.Replace(s, "?", strconv.Itoa(digit), 1)
}

func evaluateExpression(num1, num2, num3 int, operator string) bool {
	switch operator {
	case "+":
		return num1+num2 == num3
	case "-":
		return num1-num2 == num3
	case "*":
		return num1*num2 == num3
	case "/":
		return num2 != 0 && num1/num2 == num3 && num1%num2 == 0
	default:
		return false
	}
}

func findMissingDigits(equation []string) (int, int) {
	num1 := equation[0]
	num2 := equation[2]
	num3 := equation[4]
	operator := equation[1]

	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			newNum1 := replaceQuestionMark(num1, i)
			newNum3 := replaceQuestionMark(num3, j)
			val1, err1 := strconv.Atoi(newNum1)
			val3, err3 := strconv.Atoi(newNum3)

			if err1 != nil || err3 != nil {
				continue
			}

			val2, err2 := strconv.Atoi(num2)
			if err2 != nil {
				continue
			}

			if evaluateExpression(val1, val2, val3, operator) {
				return i, j
			}
		}
	}

	return -1, -1
}

func MathChallenge(str string) string {
	equation := strings.Split(str, " ")
	missing1, missing2 := findMissingDigits(equation)
	return fmt.Sprintf("%d %d", missing1, missing2)
}

func main() {
	// Test cases
	fmt.Println(MathChallenge("56? * 106 = 5?678"))   // Output: 3 9
	fmt.Println(MathChallenge("18?1 + 9 = 189?"))     // Output: 8 0
	fmt.Println(MathChallenge("38?5 * 3 = 1?595"))    // Output: 6 1
	fmt.Println(MathChallenge("50?050 / 5 = ?01010")) // Output: 5 1
}
