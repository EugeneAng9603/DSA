package main

import (
	"sort"
)

// Suffix represents a suffix of a string and its index
type Suffix struct {
	index int
	rank  [2]int
}

// BuildSuffixArray builds suffix array of given string s
func BuildSuffixArray(s string) []int {
	n := len(s)
	suffixes := make([]Suffix, n)

	// Initial ranking based on first 2 characters
	for i := 0; i < n; i++ {
		suffixes[i] = Suffix{
			index: i,
			rank:  [2]int{int(s[i]), -1},
		}
		if i+1 < n {
			suffixes[i].rank[1] = int(s[i+1])
		}
	}

	sort.Slice(suffixes, func(i, j int) bool {
		if suffixes[i].rank[0] == suffixes[j].rank[0] {
			return suffixes[i].rank[1] < suffixes[j].rank[1]
		}
		return suffixes[i].rank[0] < suffixes[j].rank[0]
	})

	ind := make([]int, n) // to get index in suffixes

	for k := 4; k < 2*n; k *= 2 {
		rank := 0
		prevRank := suffixes[0].rank
		suffixes[0].rank[0] = rank
		ind[suffixes[0].index] = 0

		for i := 1; i < n; i++ {
			if suffixes[i].rank == prevRank {
				suffixes[i].rank[0] = rank
			} else {
				prevRank = suffixes[i].rank
				rank++
				suffixes[i].rank[0] = rank
			}
			ind[suffixes[i].index] = i
		}

		for i := 0; i < n; i++ {
			nextIndex := suffixes[i].index + k/2
			if nextIndex < n {
				suffixes[i].rank[1] = suffixes[ind[nextIndex]].rank[0]
			} else {
				suffixes[i].rank[1] = -1
			}
		}

		sort.Slice(suffixes, func(i, j int) bool {
			if suffixes[i].rank[0] == suffixes[j].rank[0] {
				return suffixes[i].rank[1] < suffixes[j].rank[1]
			}
			return suffixes[i].rank[0] < suffixes[j].rank[0]
		})

	}

	suffixArr := make([]int, n)
	for i := 0; i < n; i++ {
		suffixArr[i] = suffixes[i].index
	}
	return suffixArr
}

// BuildLCP builds the Longest Common Prefix array
func BuildLCP(s string, sa []int) []int {
	n := len(s)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		rank[sa[i]] = i
	}
	k := 0
	lcp := make([]int, n)
	for i := 0; i < n; i++ {
		if rank[i] == 0 {
			k = 0
			continue
		}
		j := sa[rank[i]-1]
		for i+k < n && j+k < n && s[i+k] == s[j+k] {
			k++
		}
		lcp[rank[i]] = k
		if k > 0 {
			k--
		}
	}
	return lcp
}

// FindLongestRepeatedSubstring returns the longest substring occurring at least twice
func FindLongestRepeatedSubstringSuffix(s string) string {
	if len(s) == 0 {
		return ""
	}

	sa := BuildSuffixArray(s)
	lcp := BuildLCP(s, sa)

	maxLen := 0
	start := 0
	for i := 1; i < len(lcp); i++ {
		if lcp[i] > maxLen {
			maxLen = lcp[i]
			start = sa[i]
		}
	}

	if maxLen == 0 {
		return ""
	}
	return s[start : start+maxLen]
}
