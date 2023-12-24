// use ranks to reduce complexity of union from O(n) to O(logn)
package main

import (
	"fmt"
)

type UnionFind struct {
	parents map[int]int
	ranks   map[int]int
}

// create a union instance
func NewUnionFind() UnionFind {
	return UnionFind{
		parents: map[int]int{},
		ranks:   map[int]int{},
	}
}

// create a set in the union
func (union UnionFind) CreateSet(value int) {
	union.parents[value] = value
	union.ranks[value] = 0
}

func (union UnionFind) Find(value int) int {
	if _, found := union.parents[value]; !found {
		return -1
	}

	currParent := value
	for currParent != union.parents[currParent] {
		currParent = union.parents[currParent]
	}
	return currParent
}

func (union UnionFind) Union(valueOne, valueTwo int) {
	// check if exist
	_, parentsContainOne := union.parents[valueOne]
	_, parentsContainTwo := union.parents[valueTwo]
	if !parentsContainOne || !parentsContainTwo {
		return
	}

	valueOneRoot := union.Find(valueOne)
	valueTwoRoot := union.Find(valueTwo)

	if union.ranks[valueOneRoot] < union.ranks[valueTwoRoot] {
		union.parents[valueOneRoot] = valueTwoRoot
	} else if union.ranks[valueOneRoot] < union.ranks[valueTwoRoot] {
		union.parents[valueTwoRoot] = valueOneRoot
	} else {
		union.parents[valueTwoRoot] = valueOneRoot
		union.ranks[valueOneRoot] = union.ranks[valueOneRoot] + 1
	}
}

func main() {
	unionFind := NewUnionFind()
	fmt.Print(unionFind.Find(5))
	fmt.Print("\n")
	fmt.Print("create set 5, 10")
	unionFind.CreateSet(5)
	unionFind.CreateSet(10)

	fmt.Print("\n")
	fmt.Print(unionFind.Find(5))
	fmt.Print("\n")
	fmt.Print(unionFind.Find(10))
	fmt.Print("\n")
	fmt.Print("Do Union 5, 10")
	unionFind.Union(5, 10)
	fmt.Print("\n")
	fmt.Print(unionFind.Find(5))
	fmt.Print("\n")
	fmt.Print(unionFind.Find(10))
	fmt.Print("\n")
	fmt.Print("create Set 20")
	unionFind.CreateSet(20)
	fmt.Print("\n")
	fmt.Print(unionFind.Find(20))
	fmt.Print("\n")
	unionFind.Union(20, 10)
	fmt.Print(unionFind.Find(5))
	fmt.Print(unionFind.Find(10))
	fmt.Print(unionFind.Find(20))
}
