// *union is: {map[1:5 5:5] map[1:0 5:1]}
// union is: &{map[1:5 5:5] map[1:0 5:1]}

// use ranks to reduce complexity of union from O(n) to O(logn)
package main

// import (
//     "fmt"
// )

type UnionFind struct {
	// Write your code here.
	parents map[int]int
	ranks   map[int]int
}

func NewUnionFind() *UnionFind {
	// Write your code here.
	return &UnionFind{
		// Write your code here.
		parents: map[int]int{},
		ranks:   map[int]int{},
	}
}

// O(1), O(1)
func (union *UnionFind) CreateSet(value int) {
	// Write your code here.
	union.parents[value] = value
	union.ranks[value] = 0
}

// Time: O(a(n)), approx (O(1) Space: O(a(n)), approx (O(1))O(1)
func (union *UnionFind) Find(value int) *int {
	// Write your code here.
	if _, found := union.parents[value]; !found {
		return nil
	}

	currParent := value
	for currParent != union.parents[currParent] {
		currParent = union.parents[currParent]
	}
	//fmt.Print(currParent)
	return &currParent
}

// Time: O(a(n)), approx (O(1) Space: O(a(n)), approx (O(1))O(1)
func (union *UnionFind) Union(valueOne, valueTwo int) {
	// Write your code here.

	// check if exist
	_, parentsContainOne := union.parents[valueOne]
	_, parentsContainTwo := union.parents[valueTwo]
	if !parentsContainOne || !parentsContainTwo {
		return
	}

	valueOneRoot := *union.Find(valueOne)
	valueTwoRoot := *union.Find(valueTwo)

	if union.ranks[valueOneRoot] < union.ranks[valueTwoRoot] {
		union.parents[valueOneRoot] = valueTwoRoot
	} else if union.ranks[valueOneRoot] < union.ranks[valueTwoRoot] {
		union.parents[valueTwoRoot] = valueOneRoot
	} else {
		union.parents[valueTwoRoot] = valueOneRoot
		union.ranks[valueOneRoot] = union.ranks[valueOneRoot] + 1
	}
}
