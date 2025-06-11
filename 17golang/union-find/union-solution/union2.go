// ✔️ Solution - IV (Union and Find with Path Compression)

// There's a very small change here from Solution - II in the find() function.
// We update the par[x] as the result returned by find(par[x]) before returning from each recursive call.
// This will ensure that any future calls to find(x) won't require us to re-iterate all the way till the base par[x].
// This effectively, compresses the path to parent of x (for future calls), as the name suggests.

// Time Complexity : Θ(nlogn) , find may take O(n) for first few calls but since we are using path compression,
// each one makes subsequent searches faster and hence the amortized time complexity over a set of calls to find operation is O(logn)
// Space Complexity : O(n)

package unionsolution

// DSU with path compression (rank declared, but unused to match C++ code)
type DSUTwo struct {
	parent []int
	rank   []int
}

func NewDSUTwo(n int) *DSUTwo {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &DSUTwo{parent, rank}
}

func (dsu *DSUTwo) Find(x int) int {
	if dsu.parent[x] != x {
		dsu.parent[x] = dsu.Find(dsu.parent[x]) // path compression
	}
	return dsu.parent[x]
}

func (dsu *DSUTwo) Union(x, y int) bool {
	xp := dsu.Find(x)
	yp := dsu.Find(y)
	if xp == yp {
		return false // cycle detected
	}
	dsu.parent[xp] = yp // union without using rank (as per your code)
	return true
}

func FindRedundantConnectionUnionTwo(edges [][]int) []int {
	n := len(edges)
	dsu := NewDSUTwo(n + 1) // nodes are 1-indexed

	for _, edge := range edges {
		if !dsu.Union(edge[0], edge[1]) {
			return edge // first edge that creates a cycle
		}
	}
	return []int{}
}
