## 11. Atomic Operations

### 📘 Explanation

Atomic operations in Go are low-level, **lock-free** operations that allow safe manipulation of variables shared between goroutines. They are part of the `sync/atomic` package and are generally used with primitive types like `int32`, `int64`, `uint32`, etc.

---

### 💬 Use Case Question  
> **When would you use atomic operations instead of `sync.Mutex` in Go?**

---

### ✅ Answer:

You would use **atomic operations** when:
- You need to perform **simple, fast, lock-free operations** (like incrementing a counter).
- You’re working with a **small number of shared variables**.
- You want to avoid the **performance overhead** of acquiring and releasing a lock.

Atomic operations are ideal for:
- **Counters**
- **Flags**
- **Simple state transitions**

They are not suitable for complex state management involving multiple variables or compound data structures.

---

### 🧠 Example: Using `atomic.AddInt32` for a thread-safe counter

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int32
	var wg sync.WaitGroup

	// Launch 100 goroutines that increment the counter
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1) // Atomic increment
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
```

| Feature               | `sync/atomic`                            | `sync.Mutex`                            |
|-----------------------|------------------------------------------|------------------------------------------|
| **Use Case**          | Simple operations on primitive types     | Complex logic involving shared state     |
| **Performance**       | Very fast, lock-free                     | Slower due to locking overhead           |
| **Concurrency Safety**| Safe for basic types only                | Safe for any type of critical section    |
| **Ease of Use**       | Less intuitive, requires caution         | Easier to use for general cases          |
| **Best For**          | Counters, flags, simple shared variables | Coordinated access to complex data       |
