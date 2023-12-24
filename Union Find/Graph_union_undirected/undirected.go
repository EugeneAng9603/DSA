// find number of connected components in undirected graph
// 1->2->3  4->5 6  then there are 3 groups, so 3

package main

import "fmt"

func countConnected(graph [][]int, n int) int {
	parents := []int{}

	var i int
	for i < n {
		parents = append(parents, i)
		i += 1
	}

	ranks := []int{}
	i = 0
	for i < n {
		ranks = append(ranks, 1)
		i += 1
	}

	count := n
	for _, node := range graph {
		first := node[0]
		second := node[1]
		count -= union(first, second, parents, ranks)
		fmt.Print(node, count)
		fmt.Print("\n")
	}
	return count
}

func union(index1, index2 int, parents []int, ranks []int) int {
	p1 := find(index1, parents)
	p2 := find(index2, parents)

	// check the parents
	if p1 == p2 {
		return 0
	}
	if ranks[p1] < ranks[p2] {
		parents[p1] = p2
		ranks[p2] += ranks[p1]
	} else {
		parents[p2] = p1
		ranks[p1] += ranks[p2]
	}
	return 1
}

func find(index int, parents []int) int {
	res := index
	for res != parents[res] {
		parents[res] = parents[parents[res]]
		res = parents[res]
	}
	return res
}

func main() {
	graph := [][]int{{0, 1}, {1, 2}, {0, 2}, {3, 4}, {5, 6}, {6, 7}, {5, 7}}
	fmt.Print(countConnected(graph, 8))

}
