// Approach	Time Complexity	Space Complexity	Comment
// Brute Force	O(n * m * n)	O(1)	Easy to write, very slow for large input
// Optimized	O(m * n)	O(m * n)	Much faster for large inputs

package main

import (
	"fmt"
	"math"
)

func optimalBuilding(buildings []map[int]bool, requirements []int) (int, int) {
	n := len(buildings)
	m := len(requirements)

	// distances[requirement][i] = min distance to requirement at index i
	distances := make([][]int, m)
	for ri, req := range requirements {
		distances[ri] = make([]int, n)

		// Init with infinity
		for i := range distances[ri] {
			distances[ri][i] = math.MaxInt32
		}

		// Left to right
		lastSeen := -1
		for i := 0; i < n; i++ {
			if buildings[i][req] {
				lastSeen = i
				distances[ri][i] = 0
			} else if lastSeen != -1 {
				distances[ri][i] = i - lastSeen
			}
		}

		// Right to left
		lastSeen = -1
		for i := n - 1; i >= 0; i-- {
			if buildings[i][req] {
				lastSeen = i
			} else if lastSeen != -1 {
				distances[ri][i] = min(distances[ri][i], lastSeen-i)
			}
		}
	}

	// Find optimal building
	bestIndex := -1
	bestMaxDist := math.MaxInt32
	for i := 0; i < n; i++ {
		maxDist := 0
		for ri := 0; ri < m; ri++ {
			if distances[ri][i] > maxDist {
				maxDist = distances[ri][i]
			}
		}
		if maxDist < bestMaxDist {
			bestMaxDist = maxDist
			bestIndex = i
		}
	}

	return bestIndex, bestMaxDist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func bruteForceOptimalBuilding(buildings []map[int]bool, requirements []int) (int, int) {
	n := len(buildings)
	bestIndex := -1
	bestMaxDist := math.MaxInt32

	for i := 0; i < n; i++ {
		maxDist := 0
		for _, req := range requirements {
			minDist := math.MaxInt32
			// Search left and right
			for j := 0; j < n; j++ {
				if buildings[j][req] {
					dist := int(math.Abs(float64(i - j)))
					if dist < minDist {
						minDist = dist
					}
				}
			}
			if minDist == math.MaxInt32 {
				// requirement never met
				maxDist = math.MaxInt32
				break
			}
			if minDist > maxDist {
				maxDist = minDist
			}
		}
		if maxDist < bestMaxDist {
			bestMaxDist = maxDist
			bestIndex = i
		}
	}

	return bestIndex, bestMaxDist
}

func main() {
	buildings := []map[int]bool{
		{0: true, 1: false, 2: false},
		{0: false, 1: false, 2: false},
		{0: false, 1: true, 2: false},
		{0: false, 1: false, 2: false},
		{0: false, 1: false, 2: false},
		{0: false, 1: false, 2: true},
		{0: false, 1: false, 2: false},
		{0: true, 1: false, 2: false},
		{0: false, 1: true, 2: false},
		{0: false, 1: false, 2: true},
	}
	requirements := []int{0, 1, 2}

	idx1, dist1 := bruteForceOptimalBuilding(buildings, requirements)
	fmt.Printf("Brute Force → Index: %d, Distance: %d\n", idx1, dist1)

	idx2, dist2 := optimalBuilding(buildings, requirements)
	fmt.Printf("Optimized   → Index: %d, Distance: %d\n", idx2, dist2)
}
