// Problem:
//  Implement a fan-out (multiple goroutines processing jobs) and fan-in (gathering results from multiple goroutines) pattern in Go.

// Explanation:
// Fan-Out: Multiple workers are spawned to process jobs concurrently. Each worker takes jobs from the jobs channel and sends the results to the results channel.
// Fan-In: The main function collects results from the results channel after all workers are finished, using a sync.WaitGroup to ensure all goroutines complete before closing the results channel.
package main

import (
	"fmt"
	"sync"
)

func processJobs(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		results <- job * 2 // Simulate processing
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	var wg sync.WaitGroup
	// Fan-out: spawn multiple goroutines to process jobs
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go processJobs(jobs, results, &wg)
	}

	// Sending jobs
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	// Fan-in: collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Processed result:", result)
	}
}

// ----------------------------------------------------------------------------------------------------

// An alternative to using a sync.WaitGroup and post-sorting is to preserve order during concurrent processing by using dedicated result channels for each job.

// This method:

// Keeps job submission and result retrieval order-aligned

// Uses anonymous goroutines per job instead of a worker pool

// Avoids explicit sync.WaitGroup by using channel-based coordination
// ✅ Why This Works
// Each job has its own dedicated result channel, so even though the processing happens concurrently, you retrieve results in the same order they were launched.

// No need for sync.WaitGroup, mutex, or sorting afterward.

// It’s scalable for moderate job sizes and avoids race conditions or unordered output.

// 🚫 When Not Ideal
// Overhead of many channels and goroutines may be high for large-scale or high-frequency jobs.

// Better suited for moderate parallelism or when you care about order more than throughput.

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	type Job struct {
// 		Value    int
// 		Response chan int
// 	}

// 	jobs := make([]Job, 10)

// 	// Create jobs with individual response channels
// 	for i := 0; i < 10; i++ {
// 		resp := make(chan int)
// 		jobs[i] = Job{Value: i + 1, Response: resp}

// 		// Spawn a goroutine to process the job
// 		go func(job Job) {
// 			result := job.Value * 2 // Simulate processing
// 			job.Response <- result  // Send result to dedicated channel
// 			close(job.Response)
// 		}(jobs[i])
// 	}

// 	// Read results in the same order as jobs were submitted
// 	for _, job := range jobs {
// 		result := <-job.Response
// 		fmt.Println("Processed result:", result)
// 	}
// }
