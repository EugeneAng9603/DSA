// n, 1

package main

// input := []int{1, 2, 3, 2, 2, 1, 2}
// 	expected := 2, 2 is majority

// if same number, count++, else count--
// if found count == 0, eliminate this curr subrray, bcuz this cannot be majority
func MajorityElement(array []int) int {
	// Write your code here.
	count := 0
	var answer int

	for _, value := range array {
		if count == 0 {
			answer = value
		}

		if value == answer {
			count++
		} else {
			count--
		}
	}
	return answer
}
