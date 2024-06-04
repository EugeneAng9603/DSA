// Have the function SearchingChallenge(strArr) read the array of strings stored in
// strArr which will be a 4x4 matrix of the characters 'C', 'H', 'F', 'O',
// where C represents Charlie the dog, H represents its home, F represents dog food,
// and O represents and empty space in the grid. Your goal is to figure out
// the least amount of moves required to get Charlie to grab each piece of food in the grid
// by moving up, down, left, or right, and then make it home right after.
// Charlie cannot move onto the home before all pieces of food have been collected.
// For example: if strArr is ["FOOF", "OCOO", "OOOH", "FOOO"].

// For the input above, the least amount of steps where the dog can reach each piece of food,
// and then return home is 11 steps, so your program should return the number 11.
// The grid will always contain between 1 and 8 pieces of food.

// Examples
// Input: []string {"OOOO", "OOFF", "OCHO", "OFOO"}
// Output: 7

// Input: []string {"FOOO", "OCOH", "OFOF", "OFOO"}
// Output: 10

// Explanation:
// Directions: We define the four possible movement directions (up, down, left, right).

// Valid Check: We have a function to check if a position is within the grid bounds.

// BFS Function: We perform a breadth-first search starting from the dog's initial position.
// We use a bitmask to track which food pieces have been collected.
// We maintain a queue of states, each containing the current position,
// the collected food bitmask, and the step count. We mark states as visited to avoid revisiting them.

// Collect Food and Go Home: We add the home position as a food target
// to ensure the path includes returning home after collecting all food.

// Main Function: We parse the grid to find the initial positions of the dog,
// home, and food pieces, then call the BFS function to find the minimum steps.

// This solution ensures that the least number of steps is found by considering
// all possible paths to collect all food pieces and return home.

package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type State struct {
	Point Point
	Mask  int
	Steps int
}

func directions() []Point {
	return []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
}

func isValid(grid [][]byte, x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

func bfs(grid [][]byte, start Point, targets []Point) int {
	foodCount := len(targets)
	targetMask := (1 << foodCount) - 1 // left shift, 1 << 3 means 0001 left shift by 3 -> 1000\\\\\\\\\\\\\\\\\\\\\\\\\\

	// Find indices of food targets
	foodIndices := make(map[Point]int)
	for i, target := range targets {
		foodIndices[target] = i
	}

	// Initialize queue and visited set
	queue := []State{{start, 0, 0}}
	visited := make(map[State]bool)
	visited[State{start, 0, 0}] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		// If all food pieces have been collected and we're at home, return the steps
		if cur.Mask == targetMask && grid[cur.Point.X][cur.Point.Y] == 'H' {
			return cur.Steps
		}

		for _, d := range directions() {
			newX, newY := cur.Point.X+d.X, cur.Point.Y+d.Y
			if isValid(grid, newX, newY) && grid[newX][newY] != 'X' {
				newMask := cur.Mask
				if idx, found := foodIndices[Point{newX, newY}]; found {
					newMask |= (1 << idx)
				}
				newState := State{Point{newX, newY}, newMask, cur.Steps + 1}
				if !visited[newState] {
					visited[newState] = true
					queue = append(queue, newState)
				}
			}
		}
	}

	return -1
}

func SearchingChallenge(strArr []string) int {
	grid := make([][]byte, len(strArr))
	var dog, home Point
	var food []Point

	for i, row := range strArr {
		grid[i] = []byte(row)
		for j, char := range row {
			switch char {
			case 'C':
				dog = Point{i, j}
			case 'H':
				home = Point{i, j}
			case 'F':
				food = append(food, Point{i, j})
			}
		}
	}

	// Add home to food targets for the final path
	food = append(food, home)

	return bfs(grid, dog, food)
}

func main() {
	// Test cases
	fmt.Println(SearchingChallenge([]string{"OOOO", "OOFF", "OCHO", "OFOO"})) // Output: 7
	fmt.Println(SearchingChallenge([]string{"FOOO", "OCOH", "OFOF", "OFOO"})) // Output: 10
}
