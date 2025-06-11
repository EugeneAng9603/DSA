package main

// Problem:
//  You have a shared variable that multiple goroutines are updating concurrently.
//  If you don’t protect this variable properly, a race condition may occur.
//  Implement a solution to avoid a race condition and explain the cause of the issue.

import (
	"fmt"
	"sync"
	"time"
)

// var counter int

// func incrementCounter() {
// 	counter++
// }

// func main() {
// 	for i := 0; i < 1000; i++ {
// 		go incrementCounter()
// 	}

// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Final counter value:", counter)
// }

// Explanation of the Problem:
// There is a race condition because multiple goroutines are accessing and modifying the counter variable simultaneously.
// The counter++ operation is not atomic, so it can lead to inconsistent or incorrect results
// when multiple goroutines try to update it at the same time.

// Explanation of the Solution:
// We use a sync.Mutex (mu) to ensure that only one goroutine can modify the counter variable at a time.
// The Lock() method is called before updating the counter, and Unlock() is called after the operation is complete.
// This ensures that the counter is updated in a thread-safe manner, avoiding the race condition.

var counter int
var mu sync.Mutex

func incrementCounter() {
	mu.Lock()
	counter++
	mu.Unlock()
	// Alternatively, you can use a defer statement to ensure that the mutex is unlocked even if an error occurs.
}

func main() {
	for i := 0; i < 1000; i++ {
		go incrementCounter()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Final counter value:", counter)
}
