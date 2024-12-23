package cache

import (
	"time"
)

// https://leetcode.com/problems/lfu-cache/

type LRUCache struct {
	cacheItems   map[int]*CacheItem
	recentlyUsed LinkedList
	capacity     int
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		cacheItems:   make(map[int]*CacheItem),
		capacity:     capacity,
		recentlyUsed: NewLinkedList(),
	}
}

func (l *LRUCache) Get(key int) int {
	if c, ok := l.cacheItems[key]; ok {
		c.frequentlyCount++
		c.lastUsedTime = time.Now().UnixNano()
		l.recentlyUsed.RemoveByData(key)
		l.recentlyUsed.AddFirst(key)
		return c.value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if c, ok := l.cacheItems[key]; ok {
		c.value = value
		c.frequentlyCount++
		c.lastUsedTime = time.Now().UnixNano()

		l.recentlyUsed.RemoveByData(key)
		l.recentlyUsed.AddFirst(key)
	} else {
		if len(l.cacheItems) >= l.capacity {
			last := l.recentlyUsed.Last()
			if last != nil {
				delete(l.cacheItems, last.data)
				l.recentlyUsed.Remove(last)
			}
		}

		l.recentlyUsed.RemoveByData(key)
		l.recentlyUsed.AddFirst(key)

		l.cacheItems[key] = &CacheItem{
			key:             key,
			value:           value,
			frequentlyCount: 1,
			lastUsedTime:    time.Now().UnixNano(),
		}
	}
}
