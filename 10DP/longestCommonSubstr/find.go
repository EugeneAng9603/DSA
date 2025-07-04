package main

import (
	"find/best"
	"fmt"
	"strings"
	"time"
)

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
		got1 := best.OptimisedLongestCommonSubstring(tt.s1, tt.s2) // Using the best implementation
		duration := time.Since(start)
		status1 := "✅"
		if got1 != tt.want {
			status1 = "❌"
		}
		fmt.Printf("[%s] #%d len(s1)=%d, len(s2)=%d => got1=%d, want1=%d, time=%v\n",
			status1, i+1, len(tt.s1), len(tt.s2), got1, tt.want, duration)
	}
	for i, tt := range tests {
		start := time.Now()
		got2 := best.LongestCommonSubstring(tt.s1, tt.s2)
		duration := time.Since(start)
		status2 := "✅"
		if got2 != tt.want {
			status2 = "❌"
		}
		fmt.Printf("[%s] #%d len(s1)=%d, len(s2)=%d => got2=%d, want2=%d, time=%v\n",
			status2, i+1, len(tt.s1), len(tt.s2), got2, tt.want, duration)
	}
}
