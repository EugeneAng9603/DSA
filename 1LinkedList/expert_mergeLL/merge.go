// 2,6,7,8
// 1,3,4,5,9,10
// output is 1,2,3,4.....,10

// p1prev always follow p1
package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func MergeLinkedLists(headOne *LinkedList, headTwo *LinkedList) *LinkedList {
	// Write your code here.
	p1, p2 := headOne, headTwo
	var p1prev *LinkedList

	for p1 != nil && p2 != nil {
		// if one all smaller, just move p1prev and p1
		if p1.Value < p2.Value {
			p1prev = p1
			// no need to link p1 to p1.Next cuz it's alr linked
			p1 = p1.Next
		} else {
			// edge case if p1prev hasn't point to any
			if p1prev != nil {
				p1prev.Next = p2
			}
			// e.g. ONE is 20 larger than TWO is 1
			// p1prev move to TWO
			p1prev = p2
			// p2 should proceed to find other TWO
			p2 = p2.Next
			// link the 1 to 20
			p1prev.Next = p1
		}
	}

	// means have remaining in TWO
	// so link last of ONE to remaining TWO
	if p1 == nil {
		p1prev.Next = p2
	}

	// return the smaller linkedlist
	if headOne.Value < headTwo.Value {
		return headOne
	}
	return headTwo
}
