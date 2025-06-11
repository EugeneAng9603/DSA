## 7. How would you implement a worker pool in Go?

**Explanation:**  
A **worker pool** pattern allows you to control the number of **concurrent tasks** (workers) processing jobs.  
- Each worker listens on a **job queue** (channel),  
- Processes tasks,  
- Then waits for more jobs.

This helps to limit resource usage and control concurrency.

---

### 💬 Use Case Question:  
*What are the advantages of using a worker pool when you have a large number of jobs, but you want to limit the number of concurrent workers?*

---

### ✅ Answer:  
- **Limits concurrency:** prevents too many goroutines from running simultaneously, which could overwhelm system resources.  
- **Improves performance:** by balancing workload and efficiently reusing workers.  
- **Provides backpressure:** the job queue buffers work and workers pick them up as they become available.  
- **Simplifies management:** easier to monitor and control fixed number of workers.

---

### 💡 Example Question:  

> *"You have a task where you need to process multiple files concurrently. You want to ensure that no more than 5 files are processed at a time. Implement a worker pool using Go."*

---

### ✅ Example Solution:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Simulate file processing
func processFile(fileID int) {
    fmt.Printf("Processing file %d\n", fileID)
    time.Sleep(2 * time.Second) // Simulate time-consuming work
    fmt.Printf("Finished processing file %d\n", fileID)
}

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for fileID := range jobs {
        fmt.Printf("Worker %d started job %d\n", id, fileID)
        processFile(fileID)
        fmt.Printf("Worker %d finished job %d\n", id, fileID)
    }
}

func main() {
    const maxWorkers = 5
    const totalFiles = 10

    jobs := make(chan int, totalFiles)
    var wg sync.WaitGroup

    // Start worker pool
    for w := 1; w <= maxWorkers; w++ {
        wg.Add(1)
        go worker(w, jobs, &wg)
    }

    // Send jobs to the workers
    for fileID := 1; fileID <= totalFiles; fileID++ {
        jobs <- fileID
    }
    close(jobs) // Close channel to signal no more jobs

    wg.Wait() // Wait for all workers to finish
    fmt.Println("All files processed")
}
```