package cache

// implement cache with insert, remove, lookup operations

import (
	"fmt"
	"sync"
)

type CacheStruct struct {
	cache1 map[int](int)
	cache2 map[int](int)
	mu1    sync.RWMutex
	mu2    sync.RWMutex
}

func NewCache(c *CacheStruct) {
	c.cache1 = make(map[int]int)
	c.cache2 = make(map[int]int)
}

func Insert(c *CacheStruct, key int, value int) {
	if key%2 == 0 {
		c.mu1.Lock()
		defer c.mu1.Unlock()
		c.cache1[key] = value
	} else {
		c.mu2.Lock()
		defer c.mu2.Unlock()
		c.cache2[key] = value
	}
}

func Remove(c *CacheStruct, key int) bool {
	if key%2 == 0 {
		c.mu1.Lock()
		defer c.mu1.Unlock()
		if _, ok := c.cache1[key]; ok {
			delete(c.cache1, key)
			return true
		}
		return false
	} else {
		c.mu2.Lock()
		defer c.mu2.Unlock()
		if _, ok := c.cache2[key]; ok {
			delete(c.cache2, key)
			return true
		}
		return false
	}
}

func Lookup(c *CacheStruct, key int) bool {
	if key%2 == 0 {
		c.mu1.Lock()
		defer c.mu1.Unlock()
		if _, ok := c.cache1[key]; ok {
			return true
		}
		return false
	} else {
		c.mu2.Lock()
		defer c.mu2.Unlock()
		if _, ok := c.cache2[key]; ok {
			return true
		}
		return false
	}

}

func main() {
	var cache CacheStruct
	NewCache(&cache)
	// fmt.Println(cache)
	Insert(&cache, 1, 15)
	Insert(&cache, 2, 30)
	Insert(&cache, 4, 30)
	Insert(&cache, 6, 30)
	fmt.Println(cache.cache1, cache.cache2)

	Remove(&cache, 2)
	fmt.Println(cache.cache1, cache.cache2)

	Remove(&cache, 1)
	fmt.Println(cache.cache1, cache.cache2)

	fmt.Println(Lookup(&cache, 1))
	fmt.Println(Lookup(&cache, 6))

}
