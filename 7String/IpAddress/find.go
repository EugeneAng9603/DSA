// O(1), O(1)

// fix point 1, then proceed to try insert another 3
// by fix point 2... etc

package main

import (
	"strconv"
	"strings"
)

func ValidIPAddresses(str string) []string {
	// Write your code here.
	result := make([]string, 0)

	for i := 1; i < min(len(str), 4); i++ {
		// record result for every valid string
		currentFound := []string{"", "", "", ""}

		currentFound[0] = str[:i]
		if !checkValid(currentFound[0]) {
			continue
		}
		// min in case we reach too far
		for j := i + 1; j < i+min(len(str)-i, 4); j++ {
			currentFound[1] = str[i:j]
			if !checkValid(currentFound[1]) {
				continue
			}
			for k := j + 1; k < j+min(len(str)-j, 4); k++ {
				currentFound[2] = str[j:k]
				currentFound[3] = str[k:]
				if !checkValid(currentFound[2]) || !checkValid(currentFound[3]) {
					continue
				}

				result = append(result, strings.Join(currentFound, "."))
			}
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func checkValid(str string) bool {
	i, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	if i > 255 {
		return false
	}
	// consider if 0 or 00 or 000
	return len(str) == len(strconv.Itoa(i))
}
