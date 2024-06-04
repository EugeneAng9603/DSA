// given string, key, find the encoded string
// O(n), O(n)

package main

import "strings"

// Convert all to rune and perform %, addition then conversion to string
func CaesarCipherEncryptor(str string, key int) string {
	// Write your code here.
	shift, offset := rune(key%26), rune(26)
	runes := []rune(str)

	for i, char := range runes {
		if char >= 'a' && char+shift <= 'z' {
			char += shift
		} else {
			char += shift - offset
		}
		runes[i] = char
	}

	// RMB to convert back to string
	return string(runes)
}

// -----------------------------------------------------------------------------------------------------------------------
// solution 2 ,same complexity
// package main

// import "strings"
// // find char in alphabet = abcdefgh...
// // add it into runes arr
func CaesarCipherEncryptor1(str string, key int) string {
	// Write your code here.
	runes := []rune(str)
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for i, char := range runes {
		index := strings.Index(alphabet, string(char))
		if index == -1 {
			return ""
		}
		newIndex := (index + key) % 26
		runes[i] = rune(alphabet[newIndex])
	}
	return string(runes)
}
