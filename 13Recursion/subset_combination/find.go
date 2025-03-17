// https://leetcode.com/problems/permutations/solutions/18239/a-general-approach-to-backtracking-questions-in-java-subsets-permutations-combination-sum-palindrome-partioning/
package main

import "fmt"

func subset(nums []int) [][]int {
	result := make([][]int, 0)
	backtrackSubsets(&result, []int{}, nums, 0)
	return result
}

func backtrackSubsets(result *[][]int, subset []int, nums []int, start int) {
	newSubset := make([]int, len(subset))
	copy(newSubset, subset)
	*result = append(*result, newSubset)
	for i := start; i < len(nums); i++ {
		subset = append(subset, nums[i])
		backtrackSubsets(result, subset, nums, i+1)
		subset = subset[:len(subset)-1]
	}
}

// must not contain duplicate subsets
func subsetNoDup(nums []int) [][]int {
	result := make([][]int, 0)
	backtrackSubsetsNoDup(&result, []int{}, nums, 0)
	return result
}

func backtrackSubsetsNoDup(result *[][]int, subset []int, nums []int, start int) {
	newSubset := make([]int, len(subset))
	copy(newSubset, subset)
	*result = append(*result, newSubset)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		subset = append(subset, nums[i])
		backtrackSubsetsNoDup(result, subset, nums, i+1)
		subset = subset[:len(subset)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	backtrackCombinationSum(&result, []int{}, candidates, target, 0)
	return result
}

func backtrackCombinationSum(result *[][]int, combination []int, candidates []int, target int, start int) {
	if target == 0 {
		newCombination := make([]int, len(combination))
		copy(newCombination, combination)
		*result = append(*result, newCombination)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			continue
		}
		combination = append(combination, candidates[i])
		backtrackCombinationSum(result, combination, candidates, target-candidates[i], i)
		combination = combination[:len(combination)-1]
	}
}

func combinationSumNoDup(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	backtrackCombinationSumNoDup(&result, []int{}, candidates, target, 0)
	return result
}

func backtrackCombinationSumNoDup(result *[][]int, combination []int, candidates []int, target int, start int) {
	if target == 0 {
		newCombination := make([]int, len(combination))
		copy(newCombination, combination)
		*result = append(*result, newCombination)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			continue
		}
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		combination = append(combination, candidates[i])
		backtrackCombinationSumNoDup(result, combination, candidates, target-candidates[i], i+1)
		combination = combination[:len(combination)-1]
	}
}

// Given a string s, partition s such that every substring of the partition is a palindrome.
// Return all possible palindrome partitioning of s.
// Time  Complexity: O(N * (2 ^ N))
// Space Complexity: O(N * (2 ^ N))
func palindromePartition(s string) [][]string {
	result := make([][]string, 0)
	backtrackPalindromePartition(&result, []string{}, s, 0)
	return result
}

func backtrackPalindromePartition(result *[][]string, partition []string, s string, start int) {
	if start == len(s) {
		newPartition := make([]string, len(partition))
		copy(newPartition, partition)
		*result = append(*result, newPartition)
		return
	}
	for i := start; i < len(s); i++ {
		if isPalindrome(s[start : i+1]) {
			partition = append(partition, s[start:i+1])
			backtrackPalindromePartition(result, partition, s, i+1)
			partition = partition[:len(partition)-1]
		}
	}
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(subset([]int{1, 2, 3}))
	fmt.Print("\n")
	fmt.Println(subsetNoDup([]int{1, 2, 2}))
	fmt.Print("\n")
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Print("\n")
	fmt.Println(combinationSumNoDup([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Print("\n")
	fmt.Println(palindromePartition("aab"))

}
