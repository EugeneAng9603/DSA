// O(4^n * n), O(4^n * n)

// BACKTRACK!!
// bcuz all answer are of len 4, no need slice
// just pass path with 0000 and edit the values everytime
package main

// Global var
var table = map[byte][]byte{
	'0': {'0'},
	'1': {'1'},
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'x', 'y', 'z', 'w'},
}

func PhoneNumberMnemonics(phoneNumber string) []string {
	// Write your code here.
	// create a path with all 0 first, then change its value when reached
	path := make([]byte, len(phoneNumber))
	for i := range path {
		path[i] = '0'
	}

	results := []string{}
	helper(0, path, phoneNumber, &results)
	return results
}

func helper(idx int, path []byte, str string, results *[]string) {
	if idx == len(str) {
		// change to string
		mnemonic := string(path)
		*results = append(*results, mnemonic)
		return
	}

	digit := str[idx]
	letters := table[digit]
	for _, letter := range letters {
		path[idx] = letter
		helper(idx+1, path, str, results)
		// no need cancel anything, cuz the len is same
		// just edit it path[idx] to every possible letter
	}

}

// phoneNumber := "1905"
// 	expected := []string{
// 		"1w0j",
// 		"1w0k",
// 		"1w0l",
// 		"1x0j",
// 		"1x0k",
// 		"1x0l",
// 		"1y0j",
// 		"1y0k",
// 		"1y0l",
// 		"1z0j",
// 		"1z0k",
// 		"1z0l",
// 	}
