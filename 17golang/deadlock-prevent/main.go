// Problem:
//  Implement a Go program that demonstrates a deadlock scenario, and then fix it. Consider two goroutines trying to lock resources in the opposite order.

// Code with deadlock:
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu1 sync.Mutex
	mu2 sync.Mutex
)

// func deadlockExample1(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	mu1.Lock()
// 	fmt.Println("Goroutine 1: Acquired mu1")
// 	// Simulate time gap to increase deadlock chance
// 	time.Sleep(100 * time.Millisecond)
// 	mu2.Lock()
// 	fmt.Println("Goroutine 1: Acquired mu2")

// 	mu2.Unlock()
// 	mu1.Unlock()
// }

// func deadlockExample2(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	mu2.Lock()
// 	fmt.Println("Goroutine 2: Acquired mu2")
// 	// Simulate time gap to increase deadlock chance
// 	time.Sleep(100 * time.Millisecond)
// 	mu1.Lock()
// 	fmt.Println("Goroutine 2: Acquired mu1")

// 	mu1.Unlock()
// 	mu2.Unlock()
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go deadlockExample1(&wg)
// 	go deadlockExample2(&wg)

// 	wg.Wait() // This will hang if deadlock occurs
// }

// ❗ Why this Deadlocks:
// Goroutine 1 locks mu1 and waits for mu2.

// Goroutine 2 locks mu2 and waits for mu1.

// Neither can proceed → deadlock.

// ✅ Fix: Always Lock in Consistent Order
// Avoid deadlocks by ensuring all goroutines acquire locks in the same order:
//
// 🧠 Key Rule to Avoid Deadlocks:
// Always acquire multiple locks in the same global order (e.g., lock mu1 then mu2 in all places).

// If you must acquire multiple locks in different orders, consider using sync.RWMutex, sync.Cond, or channels to coordinate access more safely.
func noDeadlockExample(wg *sync.WaitGroup) {
	defer wg.Done()
	mu1.Lock()
	fmt.Println("Acquired mu1")
	time.Sleep(100 * time.Millisecond)
	mu2.Lock()
	fmt.Println("Acquired mu2")

	mu1.Unlock()
	mu2.Unlock()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go noDeadlockExample(&wg)
	go noDeadlockExample(&wg)
	wg.Wait() // This will not hang
}

// To prevent deadlocks, one can:
// Acquire locks in a consistent order (e.g., always acquire mu1 before mu2).
// Use timeout mechanisms like select with a time.After() channel to avoid waiting indefinitely.
// Avoid circular dependencies between goroutines that could lead to waiting on each other in an endless loop.

// ---------------------------------------
// another solution from chatgpt
// func main() {
// 	go func() {
// 		mutex1.Lock()
// 		defer mutex1.Unlock()
// 		fmt.Println("Goroutine 1: Locked mutex1")
// 		mutex2.Lock()
// 		defer mutex2.Unlock()
// 		fmt.Println("Goroutine 1: Locked mutex2")
// 	}()
// 	func() {
// 		mutex2.Lock()
// 		defer mutex2.Unlock()
// 		fmt.Println("Goroutine 2: Locked mutex2")
// 		mutex1.Lock()
// 		defer mutex1.Unlock()
// 		fmt.Println("Goroutine 2: Locked mutex1")
// 	}()
// 	// Wait for goroutines to finish
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		mutex1.Lock()
// 		defer mutex1.Unlock()
// 		fmt.Println("Goroutine 1: Finished")
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		mutex2.Lock()
// 		defer mutex2.Unlock()
// 		fmt.Println("Goroutine 2: Finished")
// 	}()
// 	wg.Wait()
// 	// Fix the deadlock by using a consistent locking order
// 	go func() {
// 		mutex1.Lock()
// 		defer mutex1.Unlock()
// 		fmt.Println("Goroutine 1: Locked mutex1")
// 		mutex2.Lock()
// 		defer mutex2.Unlock()
// 		fmt.Println("Goroutine 1: Locked mutex2")
// 	}()
// }
