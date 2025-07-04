// Time Complexity: O(m * n) bcuz O(len(s1) * len(s2))
// Space Complexity: O(n) bcuz O(len(s2))
// prev and curr are arrays (or slices) used to store
// the lengths of the longest common suffixes of substrings.
// Each has a length of len(s2) + 1 to account for the additional row/column
// that represents the empty substring (i.e., when one of the strings is empty).

// prev keeps track of the previous row of results.

// curr keeps track of the current row of results.
// After processing each row (i), the curr array is swapped with the prev array.
// This allows us to only keep track of the current and previous rows, saving memory compared to a full 2D array approach.
package best

func LongestCommonSubstring(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}

	// Use 2 rows for space optimization
	prev := make([]int, len(s2)+1)
	curr := make([]int, len(s2)+1)
	maxLen := 0

	for i := 1; i <= len(s1); i++ {
		// fmt.Printf("Processing s1[%d] = '%c'\n", i-1, s1[i-1])

		for j := 1; j <= len(s2); j++ {
			// fmt.Printf("  Comparing s1[%d] = '%c' with s2[%d] = '%c'\n", i-1, s1[i-1], j-1, s2[j-1])

			if s1[i-1] == s2[j-1] {
				curr[j] = prev[j-1] + 1
				// fmt.Printf("    Characters match! curr[%d] = %d (previous diagonal + 1)\n", j, curr[j])
				if curr[j] > maxLen {
					maxLen = curr[j]
					// fmt.Printf("    New maxLen found: %d\n", maxLen)
				}
			} else {
				curr[j] = 0
				// fmt.Printf("    Characters do not match. curr[%d] set to 0\n", j)
			}
		}

		// Swap current and previous
		prev, curr = curr, prev
		// fmt.Printf("Row %d completed. Swapped prev and curr.\n", i)
		// fmt.Printf("  prev: %v\n", prev)
		// fmt.Printf("  curr (reset): %v\n", curr)
	}

	return maxLen
}
