// Time Complexity: O(m * n)
// Space Complexity: O(n) (using two rows)

package main

import (
	"find/best"
	"fmt"
	"strings"
	"time"
)

func longestCommonSubstring(s1, s2 string) int {
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

func main() {
	tests := []struct {
		s1, s2 string
		want   int
	}{
		{"abcdfg", "zbcdfx", 4},
		{"abcdef", "xyzabc", 3},
		{"abc", "def", 0},
		{"aaaaa", "aaa", 3},
		{"", "abc", 0},
		{"abc", "", 0},
		{"abc", "abc", 3},
		{"a", "a", 1},
		{"abac", "cab", 2},
		{"abcde", "cdeab", 3},
		{"abcxyz", "xyzabc", 3},

		// 🔥 Long strings with known overlap
		{strings.Repeat("a", 100000), strings.Repeat("a", 100000), 100000},
		{strings.Repeat("ab", 50000), "z" + strings.Repeat("ab", 49999), 99998}, // offset by 1

		// 🔥 Long strings with no overlap
		{strings.Repeat("a", 100000), strings.Repeat("b", 100000), 0},

		// 🔥 Long but small overlap
		{strings.Repeat("x", 99999) + "abc", "abc" + strings.Repeat("y", 99999), 3},

		// 🔥 Large partial match
		{"hello" + strings.Repeat("x", 99995), strings.Repeat("x", 99995) + "hello", 99995},
	}

	for i, tt := range tests {
		start := time.Now()
		got := best.OptimisedLongestCommonSubstring(tt.s1, tt.s2)
		// got := best.LongestCommonSubstring(tt.s1, tt.s2) // Using the best implementation
		duration := time.Since(start)
		status := "✅"
		if got != tt.want {
			status = "❌"
		}
		fmt.Printf("[%s] #%d len(s1)=%d, len(s2)=%d => got=%d, want=%d, time=%v\n",
			status, i+1, len(tt.s1), len(tt.s2), got, tt.want, duration)
	}
}

// func main() {
// 	tests := []struct {
// 		s1, s2 string
// 		want   int
// 	}{
// 		{"abcdfg", "zbcdfx", 4}, // "bcdf"
// 		{"abcdef", "xyzabc", 3}, // "abc"
// 		{"abc", "def", 0},       // no match
// 		{"aaaaa", "aaa", 3},     // repeating chars
// 		{"", "abc", 0},          // empty
// 		{"abc", "", 0},          // empty
// 		{"abc", "abc", 3},       // same string
// 		{"a", "a", 1},           // single match
// 		{"abac", "cab", 2},      // overlapping short match
// 		{"abcde", "cdeab", 3},   // "cde"
// 		{"abcxyz", "xyzabc", 3}, // "abc"
// 	}

// 	for _, tt := range tests {
// 		got := best.OptimisedLongestCommonSubstring(tt.s1, tt.s2)
// 		if got != tt.want {
// 			fmt.Printf("❌ s1=%q s2=%q: got=%d, want=%d\n", tt.s1, tt.s2, got, tt.want)
// 		} else {
// 			fmt.Printf("✅ s1=%q s2=%q: got=%d\n", tt.s1, tt.s2, got)
// 		}
// 	}
// }
