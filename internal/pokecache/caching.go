package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mutex *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	data      []byte
}

func (c *Cache) Add(key string, data []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		data:      data,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.cache[key]; ok {
		return c.cache[key].data, true
	}
	return []byte{}, false
}

// creates new cache and starts the reapLoop go routine before returning
func NewCache(interval time.Duration) Cache {
	fmt.Println("Creating new cache..")
	cache := Cache{
		cache: map[string]cacheEntry{},
		mutex: &sync.Mutex{},
	}
	fmt.Println("Cache created. starting cache.reapLoop on a go routine...")
	go cache.reapLoop(interval)
	return cache
}

// interval is passed to the NewCache() fn – which i haven't written yet
// this func is the called by NewCache()
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			for k, v := range c.cache {
				if time.Since(v.createdAt) > interval {
					fmt.Println("reap this entry from cache --> ", k)
					delete(c.cache, k)
				}
			}
		}
	}
}
