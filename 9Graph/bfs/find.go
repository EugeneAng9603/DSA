// O(V+e), O(v)
package main

// Do not edit the class below except
// for the breadthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) BreadthFirstSearch(array []string) []string {
	// Write your code here.
	queue := make([]*Node, 0)
	queue = append(queue, n)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		array = append(array, curr.Name)
		for _, child := range curr.Children {
			queue = append(queue, child)
		}
	}

	return array
}

// var graph = NewNode("A").AddChildren("B", "C", "D")
// 	graph.Children[0].AddChildren("E").AddChildren("F")
// 	graph.Children[2].AddChildren("G").AddChildren("H")
// 	graph.Children[0].Children[1].AddChildren("I").AddChildren("J")
// 	graph.Children[2].Children[0].AddChildren("K")
// 	output := graph.BreadthFirstSearch([]string{})
// 	expected := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
// 	require.Equal(t, expected, output)
