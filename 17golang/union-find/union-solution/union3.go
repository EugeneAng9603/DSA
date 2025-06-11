// ✔️ Solution - V (Union By Rank and Find with path Compression)

// This optimization involves doing the union based on the rank (or size) of parents / representative of a component instead of just attaching one to the other randomly.
// This will ensure that even in the worst-case, we don't end up in a scenario where the par forms a skewed tree (all nodes on one-side) and we wouldn't need to iterate all nodes on each find call.
// Performing the union based on rank will keep the components/tree balanced in size.

// Union by Size Implementation
// Time Complexity : O(n*α(n)) ≈ O(n), the time complexity of each find call after union-by-rank and path compression comes out to be O(α(n)),
//  where α(n) is the inverse Ackermann function. It doesn't exceed 4 for any n < 10600 and hence is practically constant. We make O(n) calls in total.
// Space Complexity : O(n)

package unionsolution

// DSU with path compression and union by rank
type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &DSU{parent, rank}
}

// find root and assign parent to root (path compression)
// this will flatten the structure of the tree whenever we call find, making future queries faster
func (dsu *DSU) Find(x int) int {
	if dsu.parent[x] != x {
		dsu.parent[x] = dsu.Find(dsu.parent[x]) // path compression
	}
	return dsu.parent[x]
}

func (dsu *DSU) Union(x, y int) bool {
	xRoot := dsu.Find(x)
	yRoot := dsu.Find(y)
	if xRoot == yRoot {
		return false // cycle detected
	}

	// Union by rank
	if dsu.rank[xRoot] > dsu.rank[yRoot] {
		dsu.parent[yRoot] = xRoot
	} else if dsu.rank[yRoot] > dsu.rank[xRoot] {
		dsu.parent[xRoot] = yRoot
	} else {
		dsu.parent[xRoot] = yRoot
		dsu.rank[yRoot]++
	}
	return true
}

func FindRedundantConnectionUnionThree(edges [][]int) []int {
	n := len(edges)
	dsu := NewDSU(n + 1) // nodes are 1-indexed

	for _, edge := range edges {
		// if cycle is detected, return the edge
		if !dsu.Union(edge[0], edge[1]) {
			return edge
		}
	}
	return []int{}
}
