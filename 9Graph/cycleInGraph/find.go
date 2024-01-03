// O(v+e), O(v)
// 0 points to 1,3
// 1 points to 2,3,4
// check if has cycle

// detect backedge, return true
// 0 -> 1 -> 2 -> 0, this is a backedge, pointing to ancestors of node before
// for every node, check visited first
// if not, check incycle by checking if visited or in stack
// the difference is visited is for done checking for a node
// instack is when traversing from a node, track all currVisited
package main

func CycleInGraph(edges [][]int) bool {
	// Write your code here.
	numberOfNodes := len(edges)
	visited := make([]bool, numberOfNodes)
	inStack := make([]bool, numberOfNodes)

	for node := 0; node < numberOfNodes; node++ {
		if visited[node] {
			continue
		}

		inCycle := isNodeInCycle(node, edges, visited, inStack)
		if inCycle {
			return true
		}
	}
	return false
}

// if currrently in stack or detect cycle, considered true!
func isNodeInCycle(node int, edges [][]int, visited []bool, currentlyInStack []bool) bool {
	visited[node] = true
	currentlyInStack[node] = true

	// the nodes pointed by currnode
	neighbors := edges[node]
	for _, neighbor := range neighbors {
		if !visited[neighbor] {
			inCycle := isNodeInCycle(neighbor, edges, visited, currentlyInStack)
			if inCycle {
				return true
			}
		} else if currentlyInStack[neighbor] {
			return true
		}
	}
	currentlyInStack[node] = false
	return false
}

// input := [][]int{
// 	{1, 3},
// 	{2, 3, 4},
// 	{0},
// 	{},
// 	{2, 5},
// 	{},
// }
// expected := true
