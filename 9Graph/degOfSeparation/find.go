// This is an algorithm to determine which of two people is
//more socially connected within a friendship graph,
// based on the Degrees of Separation principle (popularized by the "Six Degrees of Kevin Bacon" idea).

// UNDIRECTED, UNWEIGHTED GRAPH!!

// 🧠 Core Idea
// You're given a social graph friendsLists where each person has a list of direct friends.

// For two people personOne and personTwo, compute how many other people are
// more than 6 connections away (or completely unconnected).

// The person with fewer such "distant" connections is considered more socially connected.

// ✅ Why This Algorithm Is Chosen
// Uses Breadth-First Search (BFS) to find the shortest path (degree of separation)
// from the starting person to every other node.

// BFS is optimal for unweighted graphs when finding shortest paths.

// BFS has O(V + E) time complexity — very efficient.

// By running BFS twice (once for each person), it compares their social reach.

// 📊 Time & Space Complexity
// Time Complexity: O(V + E) per findDegrees call → Total: O(V + E) * 2 = O(V + E)

// Space Complexity: O(V) for maps and queues

package main

import (
	"fmt"
	"math"
)

func degreesOfSeparation(friendsLists map[string][]string, personOne, personTwo string) string {
	degrees1 := make(map[string]int)
	visited1 := make(map[string]bool)
	findDegrees(friendsLists, personOne, degrees1, visited1)
	getUnconnectedPeople(friendsLists, degrees1)

	countOne := 0
	for _, deg := range degrees1 {
		if deg > 6 {
			countOne++
		}
	}

	fmt.Println("Degrees from", personOne, ":", degrees1)

	degrees2 := make(map[string]int)
	visited2 := make(map[string]bool)
	findDegrees(friendsLists, personTwo, degrees2, visited2)
	getUnconnectedPeople(friendsLists, degrees2)

	countTwo := 0
	for _, deg := range degrees2 {
		if deg > 6 {
			countTwo++
		}
	}

	fmt.Println("Degrees from", personTwo, ":", degrees2)

	if countOne < countTwo {
		return personOne
	} else if countTwo < countOne {
		return personTwo
	}
	return ""
}

func findDegrees(friendsLists map[string][]string, person string, degrees map[string]int, visited map[string]bool) {
	queue := [][2]interface{}{{person, 0}}
	degrees[person] = 0
	visited[person] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		currPerson := curr[0].(string)
		level := curr[1].(int)

		for _, friend := range friendsLists[currPerson] {
			if visited[friend] {
				continue
			}
			visited[friend] = true
			degrees[friend] = level + 1
			queue = append(queue, [2]interface{}{friend, level + 1})
		}
	}
}

func getUnconnectedPeople(friendsLists map[string][]string, degrees map[string]int) {
	for person := range friendsLists {
		if _, exists := degrees[person]; !exists {
			degrees[person] = math.MaxInt32
		}
	}
}

func main() {
	friendsLists := map[string][]string{
		"Alice":  {"Bob", "Claire"},
		"Bob":    {"Alice", "David"},
		"Claire": {"Alice"},
		"David":  {"Bob", "Eve"},
		"Eve":    {"David"},
		"Frank":  {}, // Disconnected person
		"George": {}, // Disconnected person
	}

	person := degreesOfSeparation(friendsLists, "Alice", "David")
	fmt.Println("More socially connected person:", person)
}
