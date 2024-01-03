// Best , logn, 1
// option 2 is logn, logn

// same, just that when found target, finalRange[0] left bound = mid
// then, finalRange[0] right bound = mid
package main

func SearchForRange(array []int, target int) []int {
	// Write your code here.
	finalRange := []int{-1, -1}
	left := 0
	right := len(array) - 1

	// first occurence, find left bound
	for left <= right {
		mid := (left + right) / 2
		if array[mid] < target {
			left = mid + 1
		} else if array[mid] > target {
			right = mid - 1
		} else {
			finalRange[0] = mid
			right = mid - 1
		}
	}

	left = 0 // no need left = 0, cuz left already stop at right + 1
	// use that to find to save effort
	right = len(array) - 1
	// find rightbound
	for left <= right {
		mid := (left + right) / 2
		if array[mid] < target {
			left = mid + 1
		} else if array[mid] > target {
			right = mid - 1
		} else {
			finalRange[1] = mid
			left = mid + 1
		}
	}
	return finalRange
}

// option 2 log, logn
// package main

// func SearchForRange(array []int, target int) []int {
// 	// Write your code here.
//     left := 0
//     right := len(array) - 1

//     for left <= right {
//         mid := (left + right) / 2
//         if array[mid] == target {
//             break
//         } else if array[mid] > target {
//             right = mid - 1
//         } else {
//             left = mid + 1
//         }
//     }

//     mid := (left + right) / 2
//     start, end := mid, mid
//     // find start
//     for array[start] == target && start > 0 {
//         start -= 1
//     }
//     // find end
//     for array[end] == target && end < len(array) - 1 {
//         end += 1
//     }

// 	return []int{start+1, end-1}
// }
