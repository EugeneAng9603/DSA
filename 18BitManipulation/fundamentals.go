package main

import "fmt"

// https://leetcode.cn/discuss/post/3571304/cong-ji-he-lun-dao-wei-yun-suan-chang-ji-enve/

func bitOperations() {
	// Example usage of the functions
	a := 0b10101010 // 170 in decimal
	b := 0b11001100 // 204 in decimal

	// Bitwise AND
	andResult := a & b // 0b10001000 (136 in decimal)

	// Bitwise OR
	orResult := a | b // 0b11101110 (238 in decimal)

	// Bitwise XOR
	xorResult := a ^ b // 0b01100110 (102 in decimal)

	// Bitwise NOT
	notA := ^a // 0b01010101 (85 in decimal)
	notB := ^b // 0b00110011 (51 in decimal)

	// Print results
	println("AND Result:", andResult)
	println("OR Result:", orResult)
	println("XOR Result:", xorResult)
	println("NOT A:", notA)
	println("NOT B:", notB)
}

// special bit operations
func bitCheckIsEven(n int) bool {
	return n&1 == 0
}
func bitCheckIsPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

// bitCount returns the number of set bits (1s) in the binary representation of n
func bitCount(n int) int {
	count := 0
	for n > 0 {
		count += n & 1 // Increment count if the least significant bit is 1
		n >>= 1        // Right shift n to process the next bit
	}
	return count
}

// bitReverse returns the binary representation of n reversed
func bitReverse(n int) int {
	reversed := 0
	for n > 0 {
		reversed <<= 1    // Shift reversed left to make space for the next bit
		reversed |= n & 1 // Add the least significant bit of n to reversed
		n >>= 1           // Right shift n to process the next bit
	}
	return reversed
}

func main() {
	bitOperations()
	fmt.Print(bitCheckIsEven(4))                      // true
	fmt.Print(bitCheckIsPowerOfTwo(8))                // true
	fmt.Println("Bit Count of 15:", bitCount(15))     // 4
	fmt.Println("Bit Reverse of 13:", bitReverse(13)) // 11 (binary 1101 reversed is 1011)
	fmt.Println("Bit Reverse of 5:", bitReverse(5))   // 5 (binary 101 reversed is 101)
}
