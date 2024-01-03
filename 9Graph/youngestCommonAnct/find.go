// O(d), O(1)
// given ancestor, find youngest common acnst
package main

type AncestralTree struct {
	Name     string
	Ancestor *AncestralTree
}

func GetYoungestCommonAncestor(top, descendantOne, descendantTwo *AncestralTree) *AncestralTree {
	// Write your code here.
	depth1 := getDepth(descendantOne, top)
	depth2 := getDepth(descendantTwo, top)

	if depth1 > depth2 {
		return backTrack(descendantOne, descendantTwo, depth1-depth2)
	} else {
		return backTrack(descendantTwo, descendantOne, depth2-depth1)
	}

}

func getDepth(node, top *AncestralTree) int {
	depth := 0
	for node != top {
		depth += 1
		node = node.Ancestor
	}
	return depth
}

func backTrack(lower, higher *AncestralTree, diff int) *AncestralTree {
	for diff != 0 {
		lower = lower.Ancestor
		diff -= 1
	}

	for lower != higher {
		lower = lower.Ancestor
		higher = higher.Ancestor
	}
	return lower
}

// func (tree *AncestralTree) addAsAncestor(descendants ...*AncestralTree) {
// 	for _, descendant := range descendants {
// 		descendant.Ancestor = tree
// 	}
// }

// func getTrees() map[rune]*AncestralTree {
// 	trees := map[rune]*AncestralTree{}
// 	for _, r := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
// 		trees[r] = &AncestralTree{Name: string(r)}
// 	}
// 	return trees
// }

// trees := getTrees()
// 	trees['A'].addAsAncestor(trees['B'], trees['C'])
// 	trees['B'].addAsAncestor(trees['D'], trees['E'])
// 	trees['D'].addAsAncestor(trees['H'], trees['I'])
// 	trees['C'].addAsAncestor(trees['F'], trees['G'])
// 	yca := GetYoungestCommonAncestor(trees['A'], trees['E'], trees['I'])
// 	require.Equal(t, trees['B'], yca)
// }
