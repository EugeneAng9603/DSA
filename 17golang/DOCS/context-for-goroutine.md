## 6. What is the context package in Go, and how is it useful in managing goroutines and cancellations?

**Explanation:**  
The `context` package in Go is used to manage the **lifecycle of goroutines**, especially for handling:  
- **Cancellation** signals  
- **Timeouts**  
- **Deadlines**  

It is commonly used for:  
- Passing **request-scoped values** in HTTP servers  
- Managing **background tasks**  
- Handling cancellation signals gracefully to avoid resource leaks and unnecessary work

---

### 💬 Use Case Question:  
*How would you use context to cancel a long-running task if the client disconnects or a timeout occurs?*

---

### 💡 Example Question:  

> *"You are implementing a web service that performs background tasks when a request is received. The task may take a long time, and you need to cancel it if the client cancels the request. How would you implement cancellation using the context package?"*
---

### ✅ Example Solution:

```go
package main

import (
    "context"
    "fmt"
    "time"
)

// longRunningTask simulates a task that respects context cancellation.
func longRunningTask(ctx context.Context) {
    for i := 0; i < 10; i++ {
        select {
        case <-ctx.Done():
            fmt.Println("Task cancelled:", ctx.Err())
            return
        default:
            fmt.Println("Working on step", i+1)
            time.Sleep(1 * time.Second) // Simulate work
        }
    }
    fmt.Println("Task completed successfully")
}

func main() {
    // Create a context that will timeout after 3 seconds
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    go longRunningTask(ctx)

    // Wait for task to complete or timeout
    time.Sleep(5 * time.Second)
    fmt.Println("Main function exiting")
}
```