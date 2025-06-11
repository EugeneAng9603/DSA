## 9. Concurrency vs Parallelism

### 📘 Explanation

- **Concurrency** is about **managing multiple tasks at the same time**, even if not running simultaneously. It’s more about structure and design of programs.
- **Parallelism** is about **executing multiple tasks simultaneously**, utilizing multiple CPU cores.

---

### 💬 Interview Follow-up:  
> **How does Go handle parallelism and concurrency differently?**

---

### ✅ Answer:

- **Concurrency in Go**:
  - Go achieves concurrency using **goroutines** and the **Go runtime scheduler**.
  - Multiple goroutines can be scheduled onto **fewer OS threads**.
  - Concurrency allows Go to manage thousands of goroutines efficiently, even on a **single CPU core**.

- **Parallelism in Go**:
  - Go supports parallelism by running goroutines across **multiple threads** on **multiple CPU cores**.
  - This is handled by the **Go scheduler**, which maps goroutines to OS threads, and those threads may run in parallel depending on core availability.

- **Go Scheduler Model**:
  - Go uses a **M:N scheduler**, meaning M goroutines are scheduled onto N threads.
  - It follows a **cooperative scheduling model**, where goroutines yield control during blocking operations (e.g., channel ops, syscalls), not via preemptive timeslices.

---

### 🧠 Summary:

| Concept       | Concurrency                                      | Parallelism                                     |
|---------------|--------------------------------------------------|--------------------------------------------------|
| **Definition**| Managing multiple tasks at once (interleaved)    | Executing multiple tasks simultaneously          |
| **Example**   | Scheduling 1000 goroutines on 2 threads           | Running 4 goroutines on 4 CPU cores              |
| **In Go**     | Achieved using goroutines and channels            | Enabled by multiple threads and CPU cores        |
| **Handled By**| Go’s M:N scheduler                                | Go runtime and underlying OS thread management   |
| **CPU Usage** | May use 1 core efficiently                       | Utilizes multiple cores for performance          |

---

Go is designed for **highly concurrent applications**, and also supports parallel execution when system resources allow it. You can control the level of parallelism using:

```go
runtime.GOMAXPROCS(numCores)
