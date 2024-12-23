package cache

import (
	"container/heap"
)

// https://leetcode.com/problems/lfu-cache/

type LFUCache struct {
	cacheItems map[int]*Item
	capacity   int
	heap       MinHeap
	clock      int
}

type MinHeap []*Item

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	if m[i].frequentlyCount == m[j].frequentlyCount {
		return m[i].lastUsedTime < m[j].lastUsedTime
	}
	return m[i].frequentlyCount < m[j].frequentlyCount
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x any) {
	item := x.(*Item)
	*m = append(*m, item)
}

func (m *MinHeap) Pop() any {
	n := len(*m)
	removed := (*m)[n-1]
	*m = (*m)[:n-1]
	return removed
}

type Item struct {
	key             int
	value           int
	frequentlyCount int
	lastUsedTime    int
}

func Constructor(capacity int) LFUCache {
	h := make(MinHeap, 0)
	heap.Init(&h)
	return LFUCache{
		heap:       h,
		cacheItems: make(map[int]*Item),
		capacity:   capacity,
		clock:      0,
	}
}

func (l *LFUCache) Get(key int) int {
	if node, ok := l.cacheItems[key]; ok {
		node.frequentlyCount++
		node.lastUsedTime = l.clock
		l.clock++
		heap.Fix(&l.heap, l.findHeapIndex(node))
		return node.value
	}
	return -1
}

func (l *LFUCache) Put(key int, value int) {
	if l.capacity == 0 {
		return
	}

	if node, ok := l.cacheItems[key]; ok {
		node.value = value
		node.frequentlyCount++
		node.lastUsedTime = l.clock
		l.clock++
		heap.Fix(&l.heap, l.findHeapIndex(node))
	} else {
		if len(l.cacheItems) >= l.capacity {
			top := heap.Pop(&l.heap).(*Item)
			delete(l.cacheItems, top.key)
		}

		newItem := &Item{
			key:             key,
			value:           value,
			frequentlyCount: 1,
			lastUsedTime:    l.clock,
		}
		l.clock++
		heap.Push(&l.heap, newItem)
		l.cacheItems[key] = newItem
	}
}

// Helper to find the heap index of an item
func (l *LFUCache) findHeapIndex(item *Item) int {
	for i, heapItem := range l.heap {
		if heapItem == item {
			return i
		}
	}
	return -1 // Should not happen if the item exists in the heap
}
