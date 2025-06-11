## 3. What are buffered and unbuffered channels, and when would you use each?

**Explanation:**  
- **Unbuffered channels** block the **sending goroutine** until another goroutine is ready to receive the value. This enforces **synchronization** between sender and receiver.
- **Buffered channels** allow a fixed number of values to be sent before blocking the sender. They decouple the sender and receiver to some extent, enabling **asynchronous communication**.

---

### 💬 Interview Follow-up: How would you choose between a buffered or unbuffered channel in different use cases?

**Example Use Case Question:**  
> *"You have a task where you need to perform multiple database writes concurrently, but you want to limit the number of concurrent database connections. How would you implement this using buffered channels?"*

---

### ✅ Answer

### 🧠 Summary

| Channel Type   | Behavior                        | Use When…                                      |
|----------------|----------------------------------|------------------------------------------------|
| **Unbuffered** | Blocks on send until receive     | You need synchronization or hand-off           |
| **Buffered**   | Allows N sends before blocking   | You need queuing or want to limit concurrency  |


**Buffered Channels:**  
Buffered channels are helpful when you want to allow **a limited number of concurrent operations** without blocking the sender immediately. In the database write scenario:
- You can use a buffered channel as a **semaphore**
- The buffer size limits the **maximum number of concurrent writes**

#### 🔧 Example:
```go
dbLimiter := make(chan struct{}, 5) // Limit to 5 concurrent DB writes

for _, record := range records {
    go func(rec Data) {
        dbLimiter <- struct{}{}        // Acquire slot
        writeToDatabase(rec)           // Perform write
        <-dbLimiter                    // Release slot
    }(record)
}
```

**Unbuffered Channels:**
Used when you want strict synchronization between senders and receivers. Each send will block until a receive is ready, and vice versa. This is useful for:
- Ensuring step-by-step processing
- Pipelining where each step depends on the previous one being completed