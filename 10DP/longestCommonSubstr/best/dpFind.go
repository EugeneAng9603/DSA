// Time Complexity: O(m * n)
// Space Complexity: O(n) (using two rows)

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
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				curr[j] = prev[j-1] + 1
				if curr[j] > maxLen {
					maxLen = curr[j]
				}
			} else {
				curr[j] = 0
			}
		}
		// Swap current and previous
		prev, curr = curr, prev
	}

	return maxLen
}
