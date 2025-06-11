// 📊 Comparison Summary Table
// Approach							Time Complexity		Space Complexity			In-place?				Notes
// Swap + Backtrack					O(n × n!)				O(n) stack				✅ Yes		Efficient and commonly used
// Backtrack + Used Track			O(n! × n)				O(n) + used				❌ No		Slower due to linear search
// Heap’s Algorithm					O(n × n!)				O(n)					✅ Yes		Fewer swaps, most efficient

package main

import "fmt"

func main() {
	fmt.Print(len(GetPermutations([]int{1, 2, 3, 4}))) // 4!=24
	fmt.Print("\n")
	fmt.Print(len(GetPermutationsNoDup([]int{1, 2, 3, 3}))) // 4!/2!1!1! = 12
	fmt.Print("\n")
	fmt.Print(GetPermutationsNoDup([]int{1, 2, 3, 3}))
}

// func main() {
// 	nums := []int{1, 2, 3}
// 	permutations := permute(nums)
// 	fmt.Println("Permutations:")
// 	for _, perm := range permutations {
// 		fmt.Println(perm)
// 	}
// 	fmt.Printf("Total permutations: %d\n", len(permutations))
// }
