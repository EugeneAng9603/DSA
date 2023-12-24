// s = ["abc", "bcd", "cbaccd"]
// output : ["b", "c"]

package main

func CommonCharacters(strings []string) []string {
	// Write your code here.
	table := map[rune]int{}
	// reduce each strings into set since ignore multiplicity
	for _, str := range strings {
		uniqueTable := map[rune]bool{}
		for _, char := range str {
			uniqueTable[char] = true
		}

		// add counter to each char in map
		// so if the table[char] = 3, means all 3 string has this char
		for char := range uniqueTable {
			table[char] += 1
		}
	}

	output := make([]string, 0)
	for char, count := range table {
		if count == len(strings) {
			output = append(output, string(char))
		}
	}
	return output
}
