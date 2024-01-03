// O(nm), O(nm), where n is width, m is height
package main

import (
	"fmt"
)

func NumberOfWaysToTraverseGraph(width int, height int) int {
	// Write your code here
	table := make([][]int, height+1)
	for i := range table {
		table[i] = make([]int, width+1)
	}

	for i := 1; i < width+1; i++ {
		for j := 1; j < height+1; j++ {
			if i == 1 || j == 1 {
				table[j][i] = 1
			} else {
				table[j][i] = table[j-1][i] + table[j][i-1]
			}
		}
	}
	fmt.Print(table)
	return table[height][width]
}

// option 2-----------factorial------------------------------------------
// O(n+m), O(1)

// Formula: ways = (xDisct + yDisct)! / xDisct! * yDisct!

// package main

// func NumberOfWaysToTraverseGraph(width int, height int) int {
// 	// Write your code here.
//     x := width - 1
//     y := height - 1

//     nume := fact(x + y)
//     deno := fact(x) * fact(y)
// 	return nume / deno
// }

// func fact(num int) int {
//     result := 1
//     for i := 2; i < num + 1; i++ {
//         result *= i
//     }
//     return result
// }
