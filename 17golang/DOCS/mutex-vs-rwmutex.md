## 5. What is the sync.Mutex and how does it differ from sync.RWMutex?

**Explanation:**  
- `sync.Mutex` provides **mutual exclusion** to protect shared resources. Only **one goroutine** can lock the mutex at a time.  
- `sync.RWMutex` provides a **reader-writer lock**:
  - Multiple **readers** can access the resource **concurrently**.
  - If a **writer** holds the lock, both readers and writers are blocked until the writer releases it.

---

### 💬 Use Case Question:  
*When would you use a `sync.RWMutex` instead of a regular `sync.Mutex`?*

---

### ✅ Answer:

- Use `sync.Mutex` if your code is mostly **writing** to a shared resource and has little to no concurrent reads. It’s simpler and more efficient because it does not involve the overhead of managing multiple readers and writers.  
- Use `sync.RWMutex` if the resource is **read-heavy** (many readers, fewer writers). It allows multiple goroutines to read **concurrently** while ensuring exclusive access to writers.

---

### 💡 Example Question:

> *"You are building a caching mechanism, where the cache is read frequently but updated infrequently. How would you implement concurrency control for this cache?"*
### ✅ Example Solution:

```go
package main

import (
    "fmt"
    "sync"
)

type Cache struct {
    data map[string]string
    mu   sync.RWMutex
}

func (c *Cache) Get(key string) (string, bool) {
    c.mu.RLock()         // Acquire read lock
    defer c.mu.RUnlock() // Release read lock
    val, ok := c.data[key]
    return val, ok
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock()         // Acquire write lock
    defer c.mu.Unlock() // Release write lock
    c.data[key] = value
}

func main() {
    cache := Cache{
        data: make(map[string]string),
    }

    var wg sync.WaitGroup

    // Simulate many readers
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            val, ok := cache.Get("foo")
            if ok {
                fmt.Printf("Reader %d got value: %s\n", id, val)
            } else {
                fmt.Printf("Reader %d found no value\n", id)
            }
        }(i)
    }

    // Simulate a writer
    wg.Add(1)
    go func() {
        defer wg.Done()
        cache.Set("foo", "bar")
        fmt.Println("Writer set value to 'bar'")
    }()

    wg.Wait()
}
```

**Explanation:**

- sync.RWMutex allows multiple goroutines to read the cache concurrently with RLock() and RUnlock().
- When writing with Lock() and Unlock(), the write is exclusive and blocks all other readers and writers to prevent race conditions.
- This design maximizes concurrency for frequent reads and safely handles infrequent writes.