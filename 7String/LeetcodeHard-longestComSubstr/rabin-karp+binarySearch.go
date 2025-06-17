// ✅ High-Level Idea:
// Use binary search on substring lengths to find the maximum length for which a duplicate substring exists.

// For each length L, use the Rabin-Karp rolling hash to check if any substring of length L appears more than once.

// Rabin-Karp enables us to compute and compare hashes of substrings efficiently, avoiding direct comparison.

// ✅ Why it's efficient:
// Binary search over length → O(log n)

// For each length, Rabin-Karp checks all substrings of that length in O(n)

// Overall time: O(n log n), space: O(n)

package main

const base = 256
const mod = 1000000007

// Check if there's a duplicate substring of length `length` using Rabin-Karp
func search(s string, length int) (int, bool) {
	n := len(s)
	if length == 0 {
		return -1, false
	}

	// Precompute base^length % mod
	var hash int
	h := 1
	for i := 0; i < length; i++ {
		h = (h * base) % mod
	}

	hashMap := make(map[int][]int)

	// Compute initial hash for first window
	for i := 0; i < length; i++ {
		hash = (hash*base + int(s[i])) % mod
	}
	hashMap[hash] = []int{0}

	for i := 1; i <= n-length; i++ {
		// Remove s[i-1], add s[i+length-1]
		hash = (hash*base - int(s[i-1])*h%mod + mod) % mod
		hash = (hash + int(s[i+length-1])) % mod

		if indices, exists := hashMap[hash]; exists {
			for _, start := range indices {
				if s[start:start+length] == s[i:i+length] {
					return start, true
				}
			}
		}
		hashMap[hash] = append(hashMap[hash], i)
	}

	return -1, false
}

// Longest duplicate substring using Rabin-Karp and binary search
func FindLongestRepeatedSubstringRabin(s string) string {
	left, right := 1, len(s)
	start := -1
	maxLen := 0

	for left <= right {
		mid := left + (right-left)/2
		if idx, found := search(s, mid); found {
			start = idx
			maxLen = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if start == -1 {
		return ""
	}
	return s[start : start+maxLen]
}
