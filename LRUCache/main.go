package main

import (
	"container/list"
	"fmt"
)

type Entry struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (lru *LRUCache) Get(key int) int {
	if elem, found := lru.cache[key]; found {
		lru.list.MoveToFront(elem)
		return elem.Value.(Entry).value
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if elem, found := lru.cache[key]; found {
		lru.list.MoveToFront(elem)
		elem.Value = Entry{key, value}
		return
	}

	if lru.list.Len() == lru.capacity {
		last := lru.list.Back()
		if last != nil {
			lru.list.Remove(last)
			delete(lru.cache, last.Value.(Entry).key)
		}
	}

	newElem := lru.list.PushFront(Entry{key, value})
	lru.cache[key] = newElem
}

func (lru *LRUCache) PrintCache() {
	for e := lru.list.Front(); e != nil; e.Next() {
		entry := e.Value.(Entry)
		fmt.Printf("[%d:%d]", entry.key, entry.value)
	}
	fmt.Println()
}

func main() {
	lru := NewLRUCache(2)

	lru.Put(1, 10)
	lru.Put(2, 20)

	fmt.Println("Get 1 :", lru.Get(1))
	// lru.PrintCache()

	lru.Put(3, 30)
	// lru.PrintCache()

	fmt.Println("Get 2 : ", lru.Get(2))
	fmt.Println("Get 3 : ", lru.Get(3))
}
