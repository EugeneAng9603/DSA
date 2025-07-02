// write a class that has an addItem method, which adds an item with an expiry.
// the class should automatically delete the item once it expires.
// Can you do it without creating multiple threads, processes, timers?
// How do you make it as real time as possible?

// solution
// implement an ExpiringItemStore in Go, without using multiple goroutines or timers per item. It uses:

// A map to store items and their expiry time.

// A min-heap (container/heap) to keep track of the next item to expire.

// Expired items are cleaned up on Add, Get, or GetAll.

// log(n)
package main

import (
	"container/heap"
	"fmt"
	"time"
)

// Item holds the value and expiry
type Item struct {
	Value     interface{}
	ExpiresAt int64
}

// ExpiryHeapItem holds key and expiry time
type ExpiryHeapItem struct {
	Key       string
	ExpiresAt int64
}

// ExpiryHeap is a min-heap of ExpiryHeapItem
type ExpiryHeap []ExpiryHeapItem

func (h ExpiryHeap) Len() int           { return len(h) }
func (h ExpiryHeap) Less(i, j int) bool { return h[i].ExpiresAt < h[j].ExpiresAt }
func (h ExpiryHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ExpiryHeap) Push(x interface{}) {
	*h = append(*h, x.(ExpiryHeapItem))
}

func (h *ExpiryHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// ExpiringItemStore is the main store
type ExpiringItemStore struct {
	items map[string]Item
	heap  ExpiryHeap
}

func NewExpiringItemStore() *ExpiringItemStore {
	h := make(ExpiryHeap, 0)
	heap.Init(&h)
	return &ExpiringItemStore{
		items: make(map[string]Item),
		heap:  h,
	}
}

// cleanup removes all expired items
func (s *ExpiringItemStore) cleanup() {
	now := time.Now().Unix()
	for s.heap.Len() > 0 {
		item := s.heap[0]
		if item.ExpiresAt > now {
			break
		}
		heap.Pop(&s.heap)
		// Only delete if the timestamp matches
		if existing, ok := s.items[item.Key]; ok && existing.ExpiresAt == item.ExpiresAt {
			delete(s.items, item.Key)
		}
	}
}

func (s *ExpiringItemStore) AddItem(key string, value interface{}, ttlSeconds int64) {
	s.cleanup()
	expiry := time.Now().Unix() + ttlSeconds
	s.items[key] = Item{Value: value, ExpiresAt: expiry}
	heap.Push(&s.heap, ExpiryHeapItem{Key: key, ExpiresAt: expiry})
}

func (s *ExpiringItemStore) GetItem(key string) (interface{}, bool) {
	s.cleanup()
	item, found := s.items[key]
	if !found {
		return nil, false
	}
	return item.Value, true
}

func (s *ExpiringItemStore) GetAllItems() map[string]interface{} {
	s.cleanup()
	result := make(map[string]interface{})
	for k, v := range s.items {
		result[k] = v.Value
	}
	return result
}

func main() {
	store := NewExpiringItemStore()
	store.AddItem("foo", "bar", 3)

	val, ok := store.GetItem("foo")
	fmt.Println("Immediate Get:", val, ok) // bar true

	time.Sleep(4 * time.Second)
	val, ok = store.GetItem("foo")
	fmt.Println("After expiry:", val, ok) // nil false
}
