package main

// Problem:
//  You need to implement a function that spawns multiple goroutines to handle jobs from a channel.
//  However, your code has a potential goroutine leak if the number of jobs is low.
//  Write the function that correctly handles this issue and explain the fix.

import (
	"fmt"
	"sync"
)

// func processJobs(jobs chan int) {
// 	for job := range jobs {
// 		fmt.Printf("Processing job: %d\n", job)
// 	}

// }

// func main() {
// 	jobs := make(chan int, 3)
// 	// Simulating jobs
// 	jobs <- 1
// 	jobs <- 2
// 	jobs <- 3

// 	// Spawning 3 goroutines to process jobs
// 	for i := 0; i < 3; i++ {
// 		go processJobs(jobs)
// 	}

// 	// Closing the channel after sending jobs
// 	close(jobs)
// }

// Problem Explanation:
// This code might cause a goroutine leak if the jobs channel is closed prematurely (before all goroutines finish processing).
// Even though the channel is closed in the main function, the goroutines might still be waiting to process jobs
// if the channel gets closed while there are still messages being processed by goroutines.

// Fix:
// You need to ensure that all goroutines finish their work before the program exits.
// This can be done by using a sync.WaitGroup to wait for all the goroutines to complete.

func processJobs(jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Processing job: %d\n", job)
	}

}
func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 3)
	// Simulating jobs
	jobs <- 1
	jobs <- 2
	jobs <- 3

	// Spawning 3 goroutines to process jobs
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go processJobs(jobs, &wg)
	}

	// Closing the channel after sending jobs
	close(jobs)
	wg.Wait()
}

// Explanation of the Fix:
// We use a sync.WaitGroup (wg) to wait for all the goroutines to finish. The Done method is called at the end of each goroutine after it processes all the jobs.
// wg.Add(1) is called before launching each goroutine to indicate the number of goroutines that need to be waited on.
// wg.Wait() ensures that the main function doesn't exit until all the goroutines finish processing.
