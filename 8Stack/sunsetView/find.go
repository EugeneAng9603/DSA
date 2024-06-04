// O(n), O(n)
package main

// given array and direction, return array of buildings which can have sunset view

// Use Stack, assume all can view
// if found a taller currBuilding, keep popping until not taller
func SunsetViews(buildings []int, direction string) []int {
	// Write your code here.
	output := make([]int, 0)
	steps := -1
	startIdx := len(buildings) - 1
	if direction == "EAST" {
		steps = 1
		startIdx = 0
	}

	idx := startIdx

	for idx >= 0 && idx < len(buildings) {
		buildingHeight := buildings[idx]

		// current stack(of height) is [21, 8,5,3,1]
		// if currHeight is 20, then keep popping all until 21 and stop
		for len(output) > 0 && buildingHeight >= buildings[output[len(output)-1]] {
			output = output[0 : len(output)-1]
		}
		output = append(output, idx)
		idx += steps

	}

	if direction == "WEST" {
		reverse(output)
	}
	return output
}

func reverse(arr []int) {
	end := len(arr) - 1
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[end-i] = arr[end-i], arr[i]
	}
}

// option 2 ------------------  runningMax ----------------
// package main

// // Use runningMax, reverse, NO NEED STACK
// func SunsetViews(buildings []int, direction string) []int {
// 	// Write your code here.
//     output := make([]int, 0)
//     steps := -1
//     startIdx := len(buildings) - 1
//     if direction == "WEST" {
//         steps = 1
//         startIdx = 0
//     }

//     idx := startIdx
//     runningMax := 0

//     for idx >= 0 && idx < len(buildings) {
//         buildingHeight := buildings[idx]

//         // if buildings is [21,8,5,3,1] only add if taller than max
//         // so [21,8]
//         if buildingHeight > runningMax {
//             output = append(output, idx)
//         }

//         runningMax = max(runningMax, buildingHeight)

//         idx += steps
//     }

//     if direction == "EAST" {
//         reverse(output)
//     }

// 	return output
// }

// func max (a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

// func reverse(arr []int){
//     end := len(arr) - 1
//     for i := 0; i < len(arr)/2; i++ {
//         arr[i], arr[end-i] = arr[end-i], arr[i]
//     }
// }
