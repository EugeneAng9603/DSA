// run length encoding, how many A, how many B, if 17A, must write 9A8A

package main

import (
	"strconv"
)

func RunLengthEncoding(str string) string {
	// Write your code here.
	output := []byte{}
	counter := 1
	// using range , char is in rune type (int32), not byte type
	// for idx, char := range str {
	//fmt.Printf("%c\n", str, reflect.TypeOf(str))
	for idx := 1; idx < len(str); idx++ {
		curr := str[idx]
		prev := str[idx-1]
		//fmt.Printf("%c, %c\n", curr, prev, reflect.TypeOf(curr))
		// fmt.Printf("Value: %c, Unicode: %U, Type: %s\n", s[i], s[i], reflect.TypeOf(s[i]))

		if curr != prev || counter == 9 {
			// strconv.Itoa(count)[0] is type byte, means getting the ASCII value of "9"
			output = append(output, strconv.Itoa(counter)[0])
			output = append(output, prev)
			counter = 0
		}
		counter++
	}

	// Last run
	output = append(output, strconv.Itoa(counter)[0])
	output = append(output, str[len(str)-1])
	return string(output)
}
