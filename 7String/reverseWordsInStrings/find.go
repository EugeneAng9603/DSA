//O(n), O(n)
// "i am great"
// output = "great am i"

package main

import (
	"strings"
)

func ReverseWordsInString(str string) string {
	// Write your code here.
	startOfWord := 0
	result := make([]string, 0)
	for idx, char := range str {
		if char == ' ' {
			currSlice := str[startOfWord:idx]
			result = append(result, currSlice)
			startOfWord = idx
		} else if str[startOfWord] == ' ' {
			result = append(result, " ")
			startOfWord = idx
		}
		// if char is a real char and startofword not ' ', continue
	}
	// final push
	result = append(result, str[startOfWord:])
	reverse(result)
	return strings.Join(result, "")
}

func reverse(list []string) {
	start := 0
	last := len(list) - 1
	for start < last {
		list[start], list[last] = list[last], list[start]
		start += 1
		last -= 1
	}
}

// --------------------option2---------------
// package main

// import (
//     "strings"
// )

// // if do this, when multiple space, we can consider
// // func ReverseWordsInString(str string) string {
// // 	// Write your code here.
// //     arr := strings.Fields(str)
// //     for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
// // 		arr[i], arr[j] = arr[j], arr[i]
// // 	}
// // 	return strings.Join(arr, " ")
// // }

// func ReverseWordsInString(str string) string {
//     startOfWord := 0
//     list := []string{}

//     for idx, char := range str {
//         if char == ' ' {
//             list = append(list, str[startOfWord:idx])
//             startOfWord = idx
//             // if last startOfWord is space and curr != space
//         } else if str[startOfWord] == ' ' {
//             list = append(list, " ")
//             startOfWord = idx
//         }
//     }
//     // final push
//     list = append(list, str[startOfWord:])

//     return strings.Join(reverse(list), "")
// }

// func reverse (list []string) []string {
//     start := 0
//     last := len(list) - 1
//     for start < last {
//         list[start], list[last] =  list[last], list[start]
//         start += 1
//         last -= 1
//     }
//     return list
// }
