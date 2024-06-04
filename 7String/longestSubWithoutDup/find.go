// O(n), min(n, a)

// Method 1:
// Record last seen for every char, if found in table && start < lastindex+1
// because if we use start = lastindex+1 for all case
// e.g. clementisacap output lementisac , should be mentisac
// becuz when reach c, start jump back to first c + 1 which is l
// but when reach second c, we alr eliminated "cle"
package main

type substring struct {
	left  int
	right int
}

func (ss substring) length() int {
	return ss.right - ss.left
}

func LongestSubstringWithoutDuplication(str string) string {
	// Write your code here.
	// last seen position of every char
	lastseen := map[byte]int{}
	longest := substring{0, 1}
	startIdx := 0
	for i, _ := range str {
		if seenIndex, found := lastseen[str[i]]; found && startIdx < seenIndex+1 { // so that we dont go back to eliminated substring
			startIdx = seenIndex + 1
		}
		if longest.length() < i+1-startIdx {
			longest = substring{startIdx, i + 1}
		}
		lastseen[str[i]] = i
	}
	return str[longest.left:longest.right]
}

// if use lastseen := map[rune]int{}, then can use lastseen[char] from i, char
// func LongestSubstringWithoutDuplication(str string) string {
// 	// Write your code here.
//     lastseen := map[rune]int{}
//     longest := substring{0, 1}
//     startIdx := 0
//     for  i, char := range str {
//         if seenIndex, found := lastseen[char]; found &&
//             startIdx < seenIndex + 1 {
//                 startIdx = seenIndex + 1
//             }
//         if longest.length() < i + 1 - startIdx {
//             longest = substring{startIdx, i + 1}
//         }
//         lastseen[char] = i
//     }
// 	return str[longest.left:longest.right]
// }

// ----------------option 2 sliding ------------------
// Template solution
// package main

// func LongestSubstringWithoutDuplication(str string) string {
//     start, end, count_repeat, longest := 0, 0, 0, 0
//     head := 0 // used to indicate where to start for the answer
//     table := make(map[byte]int)
//     for end < len(str) {
//         table[str[end]] += 1
//         if table[str[end]] > 1 {
//             count_repeat += 1
//         }
//         end += 1

//         for count_repeat > 0 {
//             if table[str[start]] > 1 {
//                 count_repeat -= 1
//             }
//             table[str[start]] -= 1
//             start += 1
//         }
//         if end - start > longest {
//             head = start
//             longest = end - start
//         }
//     }
//     return str[head:head+longest]
// }

// // func max(num1, num2 int) int {
// // 	if num1 > num2 {
// // 		return num1
// // 	}
// // 	return num2
// // }
