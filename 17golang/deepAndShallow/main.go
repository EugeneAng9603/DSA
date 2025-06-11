// Shallow Copy: A shallow copy creates a new object,
// but does not recursively copy the objects contained within the original object.
// Instead, it copies references to the objects,
// meaning both the original and copied structures point to the same underlying data.
// This can lead to issues if the data inside the copied structure is modified.

// Deep Copy: A deep copy creates a new object and recursively copies all the data,
// meaning the original and the copied structures are completely independent.
// Modifications to the copied structure won't affect the original one and vice versa.

// package main

// import "fmt"

// func main() {
// 	// Original slice
// 	original := []int{1, 2, 3, 4, 5}

// 	// Shallow copy - just referencing the original slice
// 	shallowCopy := original

// 	// Modify shallowCopy
// 	shallowCopy[0] = 100

// 	// Print both slices
// 	fmt.Println("Original slice:", original)  // [100 2 3 4 5]
// 	fmt.Println("Shallow copy:", shallowCopy) // [100 2 3 4 5]
// }

package main

import "fmt"

func main() {
	// Original slice
	original := []int{1, 2, 3, 4, 5}

	// Deep copy - creating a new slice with the same elements
	deepCopy := make([]int, len(original))
	copy(deepCopy, original)

	// Modify deepCopy
	deepCopy[0] = 100

	// Print both slices
	fmt.Println("Original slice:", original) // [1 2 3 4 5]
	fmt.Println("Deep copy:", deepCopy)      // [100 2 3 4 5]
}
