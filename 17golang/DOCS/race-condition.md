## 4. Explain the concept of the "Race Condition" in Go, and how can you prevent it?

**Explanation:**  
A **race condition** occurs when two or more goroutines access **shared resources concurrently**, and at least one of those accesses is a **write** operation. This can lead to **unpredictable and incorrect results**.

Go provides the `-race` flag to detect race conditions during testing:

```bash
go test -race
```

**Preventing Race Conditions**
You can use the following to synchronize access to shared resources and prevent race conditions:

- Mutexes (sync.Mutex)
- Channels
- Atomic operations (sync/atomic package)

Go’s runtime detects deadlocks by checking if all goroutines are blocked and unable to proceed. This is a common scenario when goroutines are waiting on locks held by each other.
To prevent deadlocks, one can:
Acquire locks in a consistent order (e.g., always acquire mu1 before mu2).
Use timeout mechanisms like select with a time.After() channel to avoid waiting indefinitely.
Avoid circular dependencies between goroutines that could lead to waiting on each other in an endless loop.

Example Use Case Question:
 "You are building a bank application, and you need to ensure that two transactions cannot modify the same account balance at the same time. How would you ensure this in Go?"

**✅ Example Solution**
To ensure that two transactions do not modify the same account balance concurrently, you can use a mutex to guard access to the shared resource:

```go

package main

import (
    "fmt"
    "sync"
)

type Account struct {
    balance int
    mu      sync.Mutex
}

func (a *Account) Deposit(amount int) {
    a.mu.Lock()
    defer a.mu.Unlock()
    a.balance += amount
}

func (a *Account) Withdraw(amount int) {
    a.mu.Lock()
    defer a.mu.Unlock()
    a.balance -= amount
}

func (a *Account) Balance() int {
    a.mu.Lock()
    defer a.mu.Unlock()
    return a.balance
}

func main() {
    acc := &Account{balance: 100}

    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        acc.Deposit(50)
        fmt.Println("Deposited 50")
    }()

    go func() {
        defer wg.Done()
        acc.Withdraw(30)
        fmt.Println("Withdrew 30")
    }()

    wg.Wait()
    fmt.Println("Final balance:", acc.Balance())
}

```