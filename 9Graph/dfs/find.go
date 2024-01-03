// O(V+e), O(v)
package main

import (
	"fmt"
)

// Do not edit the class below except
// for the depthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	// Write your code here.
	array = append(array, n.Name)
	for _, child := range n.Children {
		array = child.DepthFirstSearch(array)
		fmt.Print(array)
	}
	return array
}

// var graph = NewNode("A").AddChildren("B", "C", "D")
// 	graph.Children[0].AddChildren("E").AddChildren("F")
// 	graph.Children[2].AddChildren("G").AddChildren("H")
// 	graph.Children[0].Children[1].AddChildren("I").AddChildren("J")
// 	graph.Children[2].Children[0].AddChildren("K")
// 	output := graph.DepthFirstSearch([]string{})
// 	expected := []string{"A", "B", "E", "F", "I", "J", "C", "D", "G", "K", "H"}
// 	require.Equal(t, expected, output)
