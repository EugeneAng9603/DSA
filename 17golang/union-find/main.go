// https://leetcode.com/problems/redundant-connection/solutions/1295887/easy-solution-w-explanation-all-possible-approaches-with-comments/
package main

import (
	"fmt"
	"union/dfs"
	unionsolution "union/union-solution"
)

func main() {
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 4},
		{1, 5},
	}

	redundant1 := dfs.FindRedundantConnectionDFSone(edges)
	redundant2 := dfs.FindRedundantConnectionDFStwo(edges)
	redundant3 := unionsolution.FindRedundantConnectionUnionOne(edges)
	redundant4 := unionsolution.FindRedundantConnectionUnionTwo(edges)
	redundant5 := unionsolution.FindRedundantConnectionUnionThree(edges)
	fmt.Println("Redundant edge:", redundant1)
	fmt.Println("Redundant edge:", redundant2)
	fmt.Println("Redundant edge:", redundant3)
	fmt.Println("Redundant edge:", redundant4)
	fmt.Println("Redundant edge:", redundant5)
}

// https://leetcode.com/problems/number-of-provinces/description/
// https://leetcode.com/problems/largest-component-size-by-common-factor/description/
// https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/description/
// https://leetcode.com/problems/number-of-operations-to-make-network-connected/description/
// https://leetcode.com/problems/satisfiability-of-equality-equations/description/
// https://leetcode.com/problems/smallest-string-with-swaps/description/
// https://leetcode.com/problems/number-of-good-paths/description/
