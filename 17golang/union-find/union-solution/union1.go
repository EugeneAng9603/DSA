// ✔️ Solution - III (Union-Find)

// We can solve this using a Disjoint Subset Union (DSU). It involes two operations -

// find(x): finds the id which represents the component that a node belongs to
// union(x, y): joins the two components into a single component.
// This involves finding the representative of x-component (by find(x)) and y-component (by find(y)) and assigning them a common representative (same parent).
// We are given various edges of the graph. These can be consider as different disconnected components.
// We will join these one by one.
// A component will be represented as a tree with all its members linked to some parent and the top parent will be considered that component's representative.
// When we find an edge that joins two nodes which are already in the same component, we will return that edge as answer.
// Otherwise, we will just Union it, i.e, connect the two components by picking that edge.

// Time Complexity : O(n2), In this naive implementation, the worst case may lead to a scenario
// where we get a single component with a long lopsided chain and each find call will take O(n) for a total of n calls.
// Space Complexity : O(n), to maintain par array

package unionsolution

// DSU (Disjoint Set Union) with path compression
type DSUOne struct {
	parent []int
}

func NewDSUOne(n int) *DSUOne {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &DSUOne{parent}
}

func (dsu *DSUOne) Find(x int) int {
	if dsu.parent[x] != x {
		dsu.parent[x] = dsu.Find(dsu.parent[x]) // path compression
	}
	return dsu.parent[x]
}

func (dsu *DSUOne) Union(x, y int) bool {
	xRoot := dsu.Find(x)
	yRoot := dsu.Find(y)
	if xRoot == yRoot {
		return false // x and y are already in the same set
	}
	dsu.parent[xRoot] = yRoot
	return true
}

// Main solution using DSU
func FindRedundantConnectionUnionOne(edges [][]int) []int {
	n := len(edges)
	dsu := NewDSU(n + 1) // nodes are 1-indexed

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if !dsu.Union(u, v) {
			return edge // cycle detected
		}
	}

	return []int{}
}
