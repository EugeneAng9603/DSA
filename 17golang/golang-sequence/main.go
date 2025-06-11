// Golang program running sequence
// 💡 Note: The goroutine output order may vary due to scheduling,
// but the main goroutine's deferred statements will always follow LIFO order.
package main

import (
	"fmt"
	"time"
)

var global = initGlobal()

func initGlobal() string {
	fmt.Println("1. Global variable initialized")
	return "done"
}

func init() {
	fmt.Println("2. init() called")
}

func loop(name string) {
	defer fmt.Printf("  %s: deferred after loop ends\n", name)
	for i := 0; i < 3; i++ {
		fmt.Printf("  %s: %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	defer fmt.Println("6. deferred in main(): done message (runs just before main exits)")

	fmt.Printf("3.0 Global variable: %s\n", global)
	fmt.Println("3.1 main() started")
	defer fmt.Println("5. deferred in main(): after loop() calls")

	go loop("goroutine") // runs in new goroutine
	loop("main")         // runs in main goroutine

	time.Sleep(500 * time.Millisecond) // wait for goroutine to finish
	fmt.Printf("4. Global variable: %s\n", global)
	fmt.Println("5. main() done")

}
