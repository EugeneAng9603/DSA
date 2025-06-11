// ✔️ Solution - I (Multiple DFS Traversals)

// This solution is similar to that mentioned in the offical solution.
// We will construct the graph by adding edges one after another.
// After each addition of a new edge, we will do a dfs traversal to check if any cycle has formed.
// If a cycle is detected, we know that the last edge addition has led to the formation of cycle and hence we will simply return that edge.

// Time Complexity : O(n2), In worst case, we may need n dfs calls with each taking O(n) time complexity.
// Space Complexity : O(n), to maintain graph

package dfs

func FindRedundantConnectionDFSone(edges [][]int) []int {
	n := len(edges)
	graph := make([][]int, n+1)
	vis := make([]bool, n+1)

	for _, e := range edges {
		// Reset the visited array
		for i := range vis {
			vis[i] = false
		}

		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)

		if dfs(graph, vis, u, -1) {
			return e
		}
	}
	return []int{}
}

func dfs(graph [][]int, vis []bool, cur int, par int) bool {
	if vis[cur] {
		return true // cycle detected
	}
	vis[cur] = true

	for _, child := range graph[cur] {
		if child != par {
			if dfs(graph, vis, child, cur) {
				return true
			}
		}
	}
	return false
}
