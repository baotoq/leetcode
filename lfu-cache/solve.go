package lfu_cache

import (
	"cmp"
	"maps"
	"slices"
	"sort"
	"time"
)

// https://leetcode.com/problems/lfu-cache/

type LFUCache struct {
	cacheItems map[int]*CacheItem
	capacity   int
}

type CacheItem struct {
	key             int
	value           int
	frequentlyCount int
	lastUsedTime    int64
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cacheItems: make(map[int]*CacheItem),
		capacity:   capacity,
	}
}

func (l *LFUCache) Get(key int) int {
	if c, ok := l.cacheItems[key]; ok {
		c.frequentlyCount++
		c.lastUsedTime = time.Now().UnixNano()
		return c.value
	}
	return -1
}

func (l *LFUCache) Put(key int, value int) {
	if c, ok := l.cacheItems[key]; ok {
		c.value = value
		c.frequentlyCount++
		c.lastUsedTime = time.Now().UnixNano()
	} else {

		if len(l.cacheItems) >= l.capacity {
			keys := slices.Collect(maps.Keys(l.cacheItems))
			sort.Slice(keys, func(i, j int) bool {
				return cmp.Or(
					cmp.Compare(l.cacheItems[keys[i]].frequentlyCount, l.cacheItems[keys[j]].frequentlyCount),
					cmp.Compare(l.cacheItems[keys[i]].lastUsedTime, l.cacheItems[keys[j]].lastUsedTime),
				) < 0
			})

			delete(l.cacheItems, keys[0])
		}

		l.cacheItems[key] = &CacheItem{
			key:             key,
			value:           value,
			frequentlyCount: 1,
			lastUsedTime:    time.Now().UnixNano(),
		}
	}
}
