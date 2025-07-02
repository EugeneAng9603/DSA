
# FAANG HLD & LLD Questions with Go Solutions

## 🚀 High-Level Design (HLD) Questions

### 1. Design a URL Shortener (like Bitly)
- **Requirements**: Generate short URLs, redirect reliably, support custom aliases, high read volume.
- **Design**:
  - REST API to create and redirect URLs.
  - Use consistent hashing and sharding for scalability.
  - Store mappings in a highly available key-value DB (e.g., Redis, DynamoDB).
  - Use caching and CDN for fast redirection.

```go
type URLShortener struct {
  store map[string]string // short→long
  mu    sync.RWMutex
}

func New() *URLShortener {
  return &URLShortener{store: make(map[string]string)}
}

func (us *URLShortener) Create(short, long string) error {
  us.mu.Lock(); defer us.mu.Unlock()
  if _, ok := us.store[short]; ok {
    return errors.New("exists")
  }
  us.store[short] = long
  return nil
}

func (us *URLShortener) Get(short string) (string, bool) {
  us.mu.RLock(); defer us.mu.RUnlock()
  u, ok := us.store[short]
  return u, ok
}
```

---

### 2. Design an API Rate Limiter
- **Goal**: Protect endpoints from abuse.
- **Design**:
  - Use Redis-based token bucket or leaky bucket algorithm.
  - Handle distributed consistency across load-balanced instances.

```go
type RateLimiter struct {
  tokens    int
  capacity  int
  refillDur time.Duration
  lastRefill time.Time
  mu        sync.Mutex
}

func NewLimiter(cap, tokens int, dur time.Duration) *RateLimiter {
  return &RateLimiter{capacity: cap, tokens: tokens, refillDur: dur, lastRefill: time.Now()}
}

func (r *RateLimiter) Allow() bool {
  r.mu.Lock(); defer r.mu.Unlock()
  now := time.Now()
  elapsed := now.Sub(r.lastRefill)
  refill := int(elapsed / r.refillDur)
  if refill > 0 {
    r.tokens = min(r.capacity, r.tokens+refill)
    r.lastRefill = now
  }
  if r.tokens > 0 {
    r.tokens--
    return true
  }
  return false
}
```

---

### 3. Design a Distributed Metrics Logging System
- **Requirements**: Collect logs from thousands of servers and support rollups.
- **Design**:
  - Use partitioning, sampling, and an eventually consistent DB like Cassandra or Bigtable.

---

## 🔧 Low-Level Design (LLD) Questions

### A. Parking Lot System
- **Entities**: `ParkingLot`, `Level`, `Spot`, `Vehicle`, `Ticket`
- **Key behaviors**: spot allocation, ticket generation and payment

```go
type VehicleType int

const (
  Car VehicleType = iota
  Motorcycle
  Bus
)

type Spot struct {
  ID   string
  Type VehicleType
  Free bool
}

type Level struct {
  Spots map[string]*Spot
}

type ParkingLot struct {
  Levels []*Level
}

func (p *ParkingLot) Park(vt VehicleType) (*Ticket, error) {
  for _, lvl := range p.Levels {
    for _, sp := range lvl.Spots {
      if sp.Free && sp.Type == vt {
        sp.Free = false
        return &Ticket{SpotID: sp.ID, Entry: time.Now()}, nil
      }
    }
  }
  return nil, errors.New("no spot")
}
```

---

### B. Logging System (Singleton)

```go
type Logger struct{ /* ... */ }

var (
  logInstance *Logger
  once        sync.Once
)

func GetLogger() *Logger {
  once.Do(func() {
    logInstance = &Logger{/* setup */}
  })
  return logInstance
}
```

---

## 🧭 Approach Summary

- **HLD**: Clarify requirements, identify scaling needs and failure modes.
- **LLD**: Identify core entities, model interactions, and apply design patterns.

---

## 🗂️ Useful Resources
- [System Design on GeeksforGeeks](https://www.geeksforgeeks.org/system-design/system-design-interviews-faang/)
- [System Design Questions List (Reddit)](https://www.reddit.com/r/leetcode/comments/1j9a8u6/45_system_design_questions_i_curated_for/)
- [LLD Design Patterns Roadmap (Medium)](https://medium.com/@sandeep.kumar.ece16/low-level-design-roadmap-7581688d96fa)
