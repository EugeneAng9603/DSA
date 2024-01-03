// find highest frequency of each char in every word
// O(n * l), l is longest word

// ["this", "that", "abc"]
// output = ["t", "t", "h", "i", "s", "a", "b", "c"]
package main

func MinimumCharactersForWords(words []string) []string {
	// Write your code here.
	// overall count
	maxFreq := make(map[rune]int, 0)

	// add one in the loop for each word
	for _, word := range words {
		tempTable := make(map[rune]int, 0)
		for _, char := range word {
			tempTable[char] += 1
		}
		// update maxFreq of each word
		// like that, we dont need to check all char in maxFreq every time
		for char, count := range tempTable {
			if maxFreq[char] < count {
				maxFreq[char] = count
			}
		}
	}

	result := []string{}
	for char, count := range maxFreq {
		for count > 0 {
			result = append(result, string(char))
			count -= 1
		}
	}
	return result
}
