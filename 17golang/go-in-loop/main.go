package main

import (
	"fmt"
	"time"
)

func testloop(str string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s: %d\n", str, i)
	}
	// time.Sleep(time.Millisecond)
}
func main() {
	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Printf("%d\n", n)
		}(i)
		// time.Sleep(time.Millisecond)

	}
	fmt.Print("\n")
	// time.Sleep(time.Millisecond)
	// time.Sleep(time.Second)

	go testloop("first")
	go testloop("second")
	time.Sleep(time.Second)
}
