// "abcdcaf", first non repeating is at index 1, "b"

// one loop to check freq
// another loop to validate if freq >= 1
package main

func FirstNonRepeatingCharacter(str string) int {
	// Write your code here.
	table := map[rune]int{}

	for _, char := range str {
		table[char] += 1
	}

	for idx, char := range str {
		if table[char] == 1 {
			return idx
		}
	}

	return -1
}
