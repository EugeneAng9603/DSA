package main

var userTotal map[string]int

func init() {
	userTotal = make(map[string]int)
}

func CreateUser(cid string) {
	go func() {
		userTotal[cid]++
	}()
}

func DeleteUser(cid string) {
	go func() {
		userTotal[cid]--
	}()
}

func GetUserTotal(cid string) int {
	return userTotal[cid]
}

// —-------------------------------
// Can you mention some of the possible issues here?
// 1. Data Race (Primary Problem)
// Why?
// You are reading and writing to userTotal map concurrently from multiple goroutines without synchronization.
// go
// Copy
// Edit
// userTotal[cid]++
// userTotal[cid]--
// And:

// go
// Copy
// Edit
// return userTotal[cid]
// Result:
// Data race occurs, leading to:
// Corrupted map state.
// Possible runtime panics (fatal error: concurrent map read and map write).
// Incorrect totals (increments/decrements not applied properly).
// 2. Go Map Is Not Thread-Safe
// Built-in Go maps are not safe for concurrent access.
// This is explicitly mentioned in the Go documentation:

// It is not safe to read from and write to a map from multiple goroutines without synchronization.

// 3. Race Condition on Counter Value
// Even if map access were thread-safe, userTotal[cid]++ is not atomic:

// It does:
// Read value → increment → write back.
// Between read & write, another goroutine may read/write the same key, leading to lost updates.
// 4. Inconsistent Read in GetUserTotal
// GetUserTotal():

// go
// Copy
// Edit
// return userTotal[cid]
// Reads from userTotal without any synchronization → may read stale, partial, or inconsistent data when concurrent writes are happening.
// 5. No Way to Know When Goroutine Completes
// Both CreateUser and DeleteUser:

// go
// Copy
// Edit
// go func() {
//   userTotal[cid]++
// }()
// Launch fire-and-forget goroutines.
// No synchronization mechanism (WaitGroup, channel, etc.) to wait for completion.
// You may query GetUserTotal before goroutine actually finishes → leading to misleading results.
// 6. Unnecessary Goroutines
// Incrementing/decrementing a simple counter is fast, in-memory operation.
// Spawning goroutines introduces overhead and complicates synchronization unnecessarily.
// 🟢 Summary of Issues:
// Issue	Problem Description
// Data race on map access	Concurrent reads/writes without sync causes undefined behavior, possible panics.
// Go maps are not thread-safe	No internal locking → leads to race conditions and corruption.
// Increment/Decrement not atomic	userTotal[cid]++ and userTotal[cid]-- involve read-modify-write, not atomic.
// Inconsistent reads in GetUserTotal	No locking → may read while goroutines writing → stale/incorrect data.
// Fire-and-forget goroutines	No mechanism to ensure goroutines complete → data inconsistency risk.
// Unnecessary goroutines	For simple counter updates, goroutines are overkill and harmful.
// ✅ How to Fix?
// Option 1: Use sync.Mutex
// go
// Copy
// Edit

// var (
// 	userTotal map[string]int
// 	mu        sync.Mutex
// )

// func init() {
// 	userTotal = make(map[string]int)
// }

// func CreateUser(cid string) {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	userTotal[cid]++
// }

// func DeleteUser(cid string) {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	userTotal[cid]--
// }

// func GetUserTotal(cid string) int {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	return userTotal[cid]
// }

// Option 2: Use sync.Map (Thread-Safe Map)
// If you expect high concurrency:

// go
// Copy
// Edit

// var userTotal sync.Map

// func CreateUser(cid string) {
// 	val, _ := userTotal.LoadOrStore(cid, 0)
// 	userTotal.Store(cid, val.(int)+1)
// }

// func DeleteUser(cid string) {
// 	val, _ := userTotal.LoadOrStore(cid, 0)
// 	userTotal.Store(cid, val.(int)-1)
// }

// func GetUserTotal(cid string) int {
// 	val, ok := userTotal.Load(cid)
// 	if !ok {
// 		return 0
// 	}
// 	return val.(int)
// }

// Option 3: Use sync/atomic (If you want atomic counters)
// If only integer counters per key:

// You can design a map[string]*int64 with each counter protected by atomic operations.
// Recommendation:
// For simplicity and correctness:

// Use sync.Mutex + normal map if moderate concurrency.
// Use sync.Map if heavy concurrency or high read volume.
