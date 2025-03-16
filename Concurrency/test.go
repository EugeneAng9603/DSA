// package main

// import "fmt"

// func main() {
// 	go func() {
// 		fmt.Println("This is a goroutine!")
// 	}()
// }
// will not print cuz main returns before goroutine runs.

// -------------------------
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func() {
// 		fmt.Println("Goroutine started")
// 		// forgot to call wg.Done()
// 		// wg.Done()
// 	}()
// 	wg.Wait()
// 	fmt.Println("Main function finished")
// }

// deadlock ecause the main function waits for the goroutine to finish using wg.Wait(), but the goroutine never calls wg.Done().
// As a result, the program never exits since the main function is blocked, waiting for the goroutine to finish.

// ---------------------
package main

import (
	"fmt"
)

func main() {
	go func() {
		defer fmt.Println("This is deferred in a goroutine")
		fmt.Println("This is in the goroutine")
	}()
	defer fmt.Println("Main function")

	// time.Sleep(1 * time.Second) // adding this line works too, but not practical
}

// correct way is as follows:

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1) // Add one goroutine to wait for

// 	go func() {
// 		defer wg.Done() // Signal when the goroutine is done
// 		defer fmt.Println("This is deferred in a goroutine")
// 		fmt.Println("This is in the goroutine")
// 	}()

// 	defer fmt.Println("Main function")
// 	wg.Wait() // Wait for the goroutine to finish
// }

// will only print main function,
// --------------------
// package main

// import "fmt"

// var counter int

// func increment() {
// 	counter++
// }

// func main() {
// 	for i := 0; i < 1000; i++ {
// 		go increment()
// 	}
// 	fmt.Println("Final Counter:", counter)
// }

// the result is random, The value of counter will likely not be 1000.
// This is because the increment function modifies the shared counter variable without
// synchronization (such as using a sync.Mutex or atomic operations).
// This can lead to a race condition, where multiple goroutines concurrently modify counter,
// causing incorrect or unpredictable results.

//--------------------
