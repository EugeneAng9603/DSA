// O(n), O(1), best optimised
// solution 2 : O(n+m), O(n+m), n is len of s1, m is s2

// "hello", "hollo"
// true cuz can replace e with o
// can replace, add or remove to edit

// Bcuz all actions can only allow madeEdit once, so just keep track bool
// so if not equal, just move index accoridngly, no need slice
package main

func OneEdit(stringOne string, stringTwo string) bool {
	// Write your code here.
	len1 := len(stringOne)
	len2 := len(stringTwo)

	if abs(len1-len2) > 1 {
		return false
	}

	madeEdit := false
	index1, index2 := 0, 0
	for index1 < len1 && index2 < len2 {
		// if not equal
		if stringOne[index1] != stringTwo[index2] {
			if madeEdit {
				return false
			}
			madeEdit = true
			if len1 > len2 {
				index1 += 1
			} else if len1 < len2 {
				index2 += 1
			} else {
				index1 += 1
				index2 += 1
			}
			// if equal
		} else {
			index1 += 1
			index2 += 1
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// ------------------not optimised option---------------
// First, check length diff < 2
// Then, if found != char, check remaining string equal
// if diff length, check from i+1 == from i for the case of removing i
// O(n+m)

// This involves slicing, we can just use index so O(n) and madeEdit (bool)
// If diff char, when length equal both index ++
// If length unequal, longer's index ++
// package main

// func OneEdit(stringOne string, stringTwo string) bool {
// 	// Write your code here.
//     len1 := len(stringOne)
//     len2 := len(stringTwo)

//     if abs(len1 - len2) > 1 {
//         return false
//     }

//     for i := 0; i < min(len1, len2); i++ {
//         if stringOne[i] != stringTwo[i] {
//             if len1 > len2 {
//                 return stringOne[i+1:] == stringTwo[i:]
//             } else if len1 < len2 {
//                 return stringOne[i:] == stringTwo[i+1:]
//             } else {
//                 return stringOne[i+1:] == stringTwo[i+1:]
//             }
//         }
//     }

// 	return true
// }

// func abs (a int) int {
//     if a < 0 {
//         return -a
//     }
//     return a
// }

// func min(a, b int) int {
//     if a < b {
//         return a
//     }
//     return b
// }
