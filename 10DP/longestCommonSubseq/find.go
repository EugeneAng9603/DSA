// O(nm*min(n, m)) time | O(nm*min(n, m)) space
// DP, but keep store the longest substring in the table
// expected := []byte("XYZW")
//
//	output := LongestCommonSubsequence("ZXVVYZW", "XKYKZPW")
package main

import (
	"fmt"
)

// O(nm*min(n, m)) time | O(nm*min(n, m)) space
func LongestCommonSubsequence(s1 string, s2 string) []byte {
	// Write your code here.
	table := make([][][]byte, len(s2)+1)
	for i := range table {
		table[i] = make([][]byte, len(s1)+1)
	}

	for i := 1; i < len(table); i++ {
		for j := 1; j < len(table[i]); j++ {
			if s2[i-1] == s1[j-1] {
				tmp := make([]byte, len(table[i-1][j-1]))
				copy(tmp, table[i-1][j-1])
				table[i][j] = append(tmp, s2[i-1])
			} else {
				if len(table[i-1][j]) > len(table[i][j-1]) {
					table[i][j] = table[i-1][j]
				} else {
					table[i][j] = table[i][j-1]
				}
			}
		}
	}
	fmt.Print(table)
	return table[len(s2)][len(s1)]
}

// option 2-------------store struct----------------------------------
// O(nm) time | O(nm) space
// DP, keep track of the entry of each cell, so can reduce complexity

// package main

// type entry struct {
//     letter byte
//     length int
//     previ int
//     prevj int
// }

// func LongestCommonSubsequence(s1 string, s2 string) []byte {
// 	// Write your code here.
//     table := make([][]entry, len(s2) + 1)
//     for i := range table {
//         table[i] = make([]entry, len(s1) + 1)
//         for j := range table[i] {
//             table[i][j].letter = 0
//             table[i][j].length = 0
//             table[i][j].previ = -1
//             table[i][j].prevj = -1
//         }
//     }

//     for i := 1; i < len(s2) + 1; i++ {
//         for j := 1; j < len(s1) + 1; j++ {
//             if s2[i-1] == s1[j-1] {
//                 table[i][j] = entry{s2[i-1], table[i-1][j-1].length + 1, i-1, j-1}
//             } else {
//                 if table[i-1][j].length > table[i][j-1].length {
//                     table[i][j] = entry{0, table[i-1][j].length, i-1, j}
//                 } else {
//                     table[i][j] = entry{0, table[i][j-1].length, i, j-1}
//                 }
//             }
//         }
//     }
// 	return buildSeq(table)
// }

// func buildSeq(table [][]entry) []byte {
//     seq := make([]byte, 0)
//     i := len(table) - 1
//     j := len(table[0]) - 1

//     for i != 0 && j != 0 {
//         curr := table[i][j]
//         // means not empty string or 0 byte
//         if curr.letter != 0{
//             seq = append(seq, curr.letter)
//         }
//         i = curr.previ
//         j = curr.prevj
//     }
//     return reverse(seq)
// }

// func reverse(seq []byte) []byte {
//     for i, j := 0, len(seq)-1; i < j; i, j = i+1, j-1 {
//         seq[i], seq[j] = seq[j], seq[i]
//     }
//     return seq
// }

// option 3---------------store length--------------------------
// O(nm) time | O(nm) space
// Similar, just that only store length
// then use the logic to backtrack the string

// package main

// import (
//     "fmt"
// )

// func LongestCommonSubsequence(s1 string, s2 string) []byte {
// 	// Write your code here.
//     dp := make([][]int, len(s2) + 1)
//     for i := range dp {
//         dp[i] = make([]int, len(s1) + 1)
//     }

//     for i := 1; i < len(s2) + 1; i++ {
//         for j := 1; j < len(s1) + 1; j++ {
//             if s2[i-1] == s1[j-1] {
//                 dp[i][j] = 1 + dp[i-1][j-1]
//             } else {
//                 dp[i][j] = max(dp[i-1][j], dp[i][j-1])
//             }
//         }
//     }

//     fmt.Print(dp)
// 	return buildSeq(dp, s1)
// }

// func buildSeq(lengths [][]int, str1 string) []byte {
//     // traverse, check go up, or go left, else append, go top left
//     arr := []byte{}
//     i := len(lengths) - 1
//     j := len(lengths[0]) - 1

//     for i > 0 && j > 0 {
//         // means can go up cuz it was passed from there
//         if lengths[i][j] == lengths[i-1][j] {
//             i -= 1
//         } else if lengths[i][j] == lengths[i][j-1] {
//             j -= 1
//         } else {
//             arr = append(arr, str1[j-1])
//             i -= 1
//             j -= 1
//         }
//     }

//     // reverse here
//     for i, j := 0, len(arr) - 1; i < j; i, j = i+1, j-1 {
//         arr[i], arr[j] = arr[j], arr[i]
//     }

//     return arr
// }

// func max (a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }
