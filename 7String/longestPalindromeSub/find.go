// "abaxyzzyxxf"
// output = "xyzzyx"
// for every char from 1 to len(str), exapand left right from middle
// check both odd and even case for every char
// use getLongest sub-function

// O(n2), O(n)
package main

type substring struct {
	left  int
	right int
}

func (ss substring) length() int {
	return ss.right - ss.left
}

func LongestPalindromicSubstring(str string) string {
	// Write your code here.
	result := substring{0, 1}
	for i := 1; i < len(str); i++ {
		odd := getLongest(str, i-1, i+1)
		even := getLongest(str, i-1, i)
		longest := even
		// compare even to odd
		if odd.length() > even.length() {
			longest = odd
		}
		// compare curr to longest
		if longest.length() > result.length() {
			result = longest
		}
	}
	return str[result.left:result.right]
}

func getLongest(str string, left, right int) substring {
	// to edge of string
	for left >= 0 && right < len(str) {
		if str[left] != str[right] {
			break
		}
		left -= 1
		right += 1
	}
	// bcuz we already reached beyond left and right, so should use
	// left+1 and right-1 , and to include right-1, we use right
	return substring{left + 1, right}
}
