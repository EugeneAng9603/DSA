// ✔️ Solution - II (Single DFS Traversal)

// We need to find a cycle in the given graph and return an edge in that cycle that occurs last in input edges.
// We are given that a cycle always exist in the given graph and hence an answer will always exist.
// So, we only need to find the cycle and the edges involved in that cycle of the graph.
// We will store these edges in a set called cycle.

// For this we can use DFS traversal with small modification.
// We will start the dfs on the given graph from node 1 and maintain a vis array denoting the nodes that are visited.
// We will also keep a par variable to denote the parent of current node cur. In each dfs. we will iterate over the child nodes of cur.
// If we ever visit an already visited node, we know that we are in a cycle. So, we will mark this node as the cycleStart and return.
// We will start pushing all nodes from the recursion stack until we reach cycleStart node (its first visit recursive call)

// We have the following cases:

// child == par : The child node of cur is it's parent. So, skip it.
// cycle.empty() == true: We haven't yet detected any cycle. So we will continue exploring the graph with recursive dfs calls.
// cycleStart != -1: We have detected a cycle and marked the starting node of it. We will stop further dfs calls as cycle is found. We push all the nodes till we reach back to node numbered - cycleStart since they are all part of the cycle.
// cur == cycleStart: We have reached back to the start of cycle. By now, we have pushed all nodes in the cycle. So, just mark cycleStart as -1 denoting we don't need to further push any nodes and return.
// Finally, cycle contains all the nodes of cycle in the graph.
// We will iterate over the input edges in the reverse order to find the last edge in it that's part of cycle.
// If both node of an edge is in cycle, then we will return that edge.

// Alternate Implementation (Faster Runtime ~4ms)
// Time Complexity : O(n), where n is the number of nodes in the graph. We need O(n) for dfs and another O(n) for iterating through the input for finding the last edge in it that occurs in the cycle.
// Space Complexity : O(n), for maintaining graph and cycle

package dfs

var (
	cycle      = make(map[int]bool) // holds all nodes that are part of a cycle
	cycleStart = -1                 // node where the cycle was first detected
)

func FindRedundantConnectionDFStwo(edges [][]int) []int {
	n := len(edges)
	graph := make([][]int, n+1)
	vis := make([]bool, n+1)

	// Build the undirected graph
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// Start DFS to detect the cycle
	dfsTwo(graph, vis, 1, -1)

	// Walk through the edges in reverse order to find the last one closing the cycle
	for i := n - 1; i >= 0; i-- {
		u, v := edges[i][0], edges[i][1]
		if cycle[u] && cycle[v] {
			return edges[i]
		}
	}

	return []int{}
}

func dfsTwo(graph [][]int, vis []bool, cur int, par int) {
	if vis[cur] {
		cycleStart = cur
		return
	}

	vis[cur] = true

	for _, child := range graph[cur] {
		if child == par {
			continue
		}
		if len(cycle) == 0 {
			dfsTwo(graph, vis, child, cur)
		}
		if cycleStart != -1 {
			cycle[cur] = true
		}
		if cur == cycleStart {
			cycleStart = -1
			return
		}
	}
}
