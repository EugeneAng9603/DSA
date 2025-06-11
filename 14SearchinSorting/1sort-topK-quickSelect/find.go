// Best, average: O(n), O(1)
// Worst: O(n^2), O(1)

// k smallest of that array
// quick select: the last value will be corrected first
// followed by second last
// so quick select first, after done for one pivot, check if
// the "right" position is on the kth smallest place, if yes that's the answer
// bcuz only this 'right' is at the correct place

// expected := 5
// output := Quickselect([]int{8, 5, 2, 9, 7, 6, 3}, 3)

// randomised array so on average the time complexity is O(n)
package main

func Quickselect(array []int, k int) int {
	// Write your code here.
	return helper(array, 0, len(array)-1, k-1)
}

func helper(arr []int, start, end, position int) int {
	for {
		if start > end {
			panic("this should be wrong")
		}
		pivot, left, right := start, start+1, end
		for left <= right {
			// if both > right, found case, swap so smallest go in front
			if arr[left] > arr[right] && arr[right] < arr[pivot] {
				swap(left, right, arr)
			}
			// left is smaller, check next val
			if arr[left] <= arr[pivot] {
				left += 1
			}
			// right is larger, check next val
			if arr[right] >= arr[pivot] {
				right -= 1
			}
		}
		// done for one pivot, make pivot to the correct position
		swap(pivot, right, arr)

		// now check if done for kth smallest, if not continue
		// since 'right' is on the correct position
		// by doing this just target to sort for the half
		if right == position {
			return arr[right]
			// do values to the right
		} else if right < position {
			start = right + 1
			// do values to the left
		} else {
			end = right - 1
		}
	}
}

func swap(one, two int, arr []int) {
	arr[one], arr[two] = arr[two], arr[one]
}
