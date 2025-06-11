// Problem:
//  Implement a rate-limited counter that allows no more than 5 operations per second, and prints "Rate limit exceeded" when the limit is reached.

// Explanation:
// We create a ticker to reset the counter every second. The rateLimiter function limits operations to 5 per second by resetting the counter after each tick.
// If the rate limit is exceeded, it prints "Rate limit exceeded".
// The select statement is used to wait for either the ticker to reset the count or for a new operation to come in.

// Operation successful
// Operation successful
// Operation successful
// Operation successful
// Operation successful
// Rate limit exceeded
// Rate limit exceeded
// Rate limit exceeded
// Rate limit exceeded
// Rate limit exceeded
// Rate limit exceeded
// Rate limit exceeded
// Rate limit exceeded
// Rate limit exceeded
// Operation successful
// Operation successful
// Operation successful
// Operation successful
// Operation successful
// Rate limit exceeded
// Rate limit exceeded
// Rate limit exceeded
// Rate limit exceeded
// ....
package main

import (
	"fmt"
	"time"
)

func rateLimiter(rateLimit int) {
	// This creates a ticker that sends an event every 1 second. It's used to reset the operation counter (count) each second.
	ticker := time.NewTicker(time.Second)
	//Ensures the ticker is properly stopped when the function exits (good resource management practice).
	defer ticker.Stop()
	count := 0

	for {
		select {
		// Reset count every second via the ticker.
		case <-ticker.C:
			count = 0
			// Every 50 milliseconds, it attempts an operation:
		case <-time.After(50 * time.Millisecond): // Simulating an operation attempt
			if count < rateLimit {
				count++
				fmt.Println("Operation successful")
			} else {
				fmt.Println("Rate limit exceeded")
			}
		}
	}
}

func main() {
	go rateLimiter(5) // Allow 5 operations per second
	time.Sleep(10 * time.Second)
}
