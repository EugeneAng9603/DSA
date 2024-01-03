package main

// Do not edit the class below except for the insertKeyValuePair,
// getValueFromKey, and getMostRecentKey methods. Feel free
// to add new properties and methods to the class.
type LRUCache struct {
	maxSize int
	// Add fields here.
	index            map[string]*Node
	currSize         int
	listOfMostRecent *DoublyLinkedList
}

type Node struct {
	key   string
	value int
	prev  *Node
	next  *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func NewLRUCache(size int) *LRUCache {
	// Write your code here.
	lru := &LRUCache{
		maxSize:          size,
		index:            map[string]*Node{},
		currSize:         0,
		listOfMostRecent: &DoublyLinkedList{},
	}
	if lru.maxSize < 1 {
		lru.maxSize = 1
	}
	return lru
}

// 1,1
func (cache *LRUCache) InsertKeyValuePair(key string, value int) {
	// Write your code here.

	// check if alredy in cache, if not found
	if _, found := cache.index[key]; !found {
		// check size
		if cache.currSize == cache.maxSize {
			cache.evictLeastRecent()
		} else {
			cache.currSize += 1
		}
		// add to hashmap
		cache.index[key] = &Node{
			key:   key,
			value: value,
		}
		// if alr in, just update value
	} else {
		cache.replaceKey(key, value)
	}

	// always update recent
	cache.updateMostRecent(cache.index[key])
}

// 1,1
// The second return value indicates whether or not the key was found
// in the cache.
func (cache *LRUCache) GetValueFromKey(key string) (int, bool) {
	// Write your code here.
	if node, found := cache.index[key]; !found {
		return 0, false
	} else {
		cache.updateMostRecent(node)
		return node.value, true
	}

}

// 1,1
// The second return value is false if the cache is empty.
func (cache *LRUCache) GetMostRecentKey() (string, bool) {
	// Write your code here.
	if cache.listOfMostRecent.head == nil {
		return "", false
	}
	return cache.listOfMostRecent.head.key, true
}

func (cache *LRUCache) evictLeastRecent() {
	key := cache.listOfMostRecent.tail.key
	// remove from linked list
	cache.listOfMostRecent.removeTail()
	// delete from hashmap
	delete(cache.index, key)
}

func (cache *LRUCache) updateMostRecent(node *Node) {
	cache.listOfMostRecent.setHeadTo(node)
}

// just change value, for linkedlist handle in updateMostRecent
func (cache *LRUCache) replaceKey(key string, value int) {
	if node, found := cache.index[key]; !found {
		panic("The provided key isn't in the cache!")
	} else {
		node.value = value
	}
}

func (list *DoublyLinkedList) setHeadTo(node *Node) {
	// already head
	if list.head == node {
		return
	}
	// empty list
	if list.head == nil {
		list.head, list.tail = node, node
		return
	}
	// only one node
	if list.head == list.tail {
		list.tail.prev = node
		list.head = node
		list.head.next = list.tail
		return
	}

	// normal case, not head not empty, not only one node
	if list.tail == node {
		list.removeTail()
	}

	// update
	node.removeBindings()
	list.head.prev = node
	node.next = list.head
	list.head = node
}

func (list *DoublyLinkedList) removeTail() {
	// empty
	if list.tail == nil {
		return
	}
	// only one node
	if list.tail == list.head {
		list.head, list.tail = nil, nil
		return
	}
	// normal case
	list.tail = list.tail.prev
	list.tail.next = nil
}

func (node *Node) removeBindings() {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev, node.next = nil, nil
}

// -------------------------------------------------------------------------------------
// This file is initialized with a code version of this
// question's sample test case. Feel free to add, edit,
// or remove test cases in this file as you see fit!

// func (s *TestSuite) TestCase1(t *TestCase) {
// 	lruCache := NewLRUCache(3)
// 	lruCache.InsertKeyValuePair("b", 2)
// 	lruCache.InsertKeyValuePair("a", 1)
// 	lruCache.InsertKeyValuePair("c", 3)
// 	key, _ := lruCache.GetMostRecentKey()
// 	require.True(t, key == "c")
// 	value, _ := lruCache.GetValueFromKey("a")
// 	require.True(t, value == 1)
// 	key, _ = lruCache.GetMostRecentKey()
// 	require.True(t, key == "a")
// 	lruCache.InsertKeyValuePair("d", 4)
// 	_, found := lruCache.GetValueFromKey("b")
// 	require.True(t, !found)
// 	lruCache.InsertKeyValuePair("a", 5)
// 	value, _ = lruCache.GetValueFromKey("a")
// 	require.True(t, value == 5)
// }
