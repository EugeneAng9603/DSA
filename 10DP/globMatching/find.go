// nm, nm

// * wildcards, anything any length
// ? any single char
// if reached *, check left || top
// if reached ? or equal, dp[i][j] = dp[i-1][j-1]
// else must be false (default value in Go)

// fileName := "abcdefg"
// pattern := "a*e?g"
// output := GlobMatching(fileName, pattern)
// require.True(t, output)

package main

func GlobMatching(fileName, pattern string) bool {
	// Write your code here.
	table := make([][]bool, len(fileName)+1)
	for i := 0; i < len(fileName)+1; i++ {
		table[i] = make([]bool, len(pattern)+1)
	}
	table[0][0] = true
	// edge case handling
	// if start with * or ******
	for j := 1; j < len(pattern)+1; j++ {
		if pattern[j-1] == '*' {
			table[0][j] = table[0][j-1]
		}
	}

	for i := 1; i < len(fileName)+1; i++ {
		for j := 1; j < len(pattern)+1; j++ {
			if pattern[j-1] == '*' {
				table[i][j] = table[i][j-1] || table[i-1][j]
			} else if pattern[j-1] == '?' || pattern[j-1] == fileName[i-1] {
				table[i][j] = table[i-1][j-1]
			}
		}
	}
	return table[len(fileName)][len(pattern)]
}
