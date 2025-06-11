# Go Interview Question: Goroutines vs Threads

## 1. What is the difference between a goroutine and a thread in Go?

**Explanation:**  
Go's concurrency model is based on **goroutines**, which are lightweight threads managed by the Go runtime. Threads, on the other hand, are managed by the operating system.

A **goroutine** is:
- Cheaper in terms of memory and scheduling
- Managed by the Go runtime, not the OS
- Multiplexed onto a smaller number of OS threads by the Go scheduler

---

### Interview Follow-up: How does Go handle scheduling and blocking of goroutines?

**Answer:**  
Go uses an **M:N scheduler** to manage goroutines. This means:
- Multiple goroutines (`M`) are scheduled onto fewer OS threads (`N`)
- The Go runtime uses **goroutine queues** to manage execution

#### Key behaviors:
- Goroutines are non-blocking by default
- If a goroutine blocks (e.g., waiting on a channel or a lock), the scheduler **suspends it** and switches to another ready goroutine
- This enables Go to fully utilize available **CPU cores**

#### Practical Tool:
- `sync.WaitGroup` is crucial for managing goroutines
  - It allows the `main` function to wait for all goroutines to finish
