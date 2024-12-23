package cache

// https://leetcode.com/problems/lfu-cache/

type LRUCache struct {
	cacheItems   map[int]*Node
	recentlyUsed LinkedList
	capacity     int
}

// Constructor initializes the LRU Cache with a given capacity
func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		cacheItems:   make(map[int]*Node),
		capacity:     capacity,
		recentlyUsed: NewLinkedList(),
	}
}

// Get retrieves a value from the cache and updates its position
func (l *LRUCache) Get(key int) int {
	if node, ok := l.cacheItems[key]; ok {
		// Move the accessed node to the front
		l.recentlyUsed.Remove(node)
		l.recentlyUsed.AddFirst(node)
		return node.value
	}
	return -1
}

// Put inserts or updates a key-value pair in the cache
func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.cacheItems[key]; ok {
		// Update the value and move node to the front
		node.value = value
		l.recentlyUsed.Remove(node)
		l.recentlyUsed.AddFirst(node)
	} else {
		// Evict the least recently used node if capacity is reached
		if len(l.cacheItems) >= l.capacity {
			last := l.recentlyUsed.Last()
			if last != nil {
				delete(l.cacheItems, last.key)
				l.recentlyUsed.Remove(last)
			}
		}

		// Add the new node
		node := &Node{key: key, value: value}
		l.recentlyUsed.AddFirst(node)
		l.cacheItems[key] = node
	}
}
