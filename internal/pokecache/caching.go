package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mutex *sync.RWMutex
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
	fmt.Printf("\n\n%v\n\n", "ADDING ENTRY TO CACHE")
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.cache[key]; ok {
		fmt.Printf("\n\n%v\n\n", "USING CACHE")
		return c.cache[key].data, true
	}
	return []byte{}, false
}

// creates new cache and starts the reapLoop go routine before returning
func NewCache(interval time.Duration) *Cache {
	fmt.Println("Creating new cache..")
	cache := &Cache{
		cache: map[string]cacheEntry{},
		mutex: &sync.RWMutex{},
	}

	fmt.Println("Cache created. starting cache.reapLoop on a go routine...")
	go cache.reapLoop(interval)
	fmt.Println("finished NewCache without error")
	return cache
}

// interval is passed to the NewCache() fn – which i haven't written yet
// this func is the called by NewCache()
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	// missing a done channel to stop go routine
	for {
		select {
		case <-ticker.C:
			// mutex locks need to go AROUND the for loop for obvious reasons
			c.mutex.Lock()
			for k, v := range c.cache {
				if time.Since(v.createdAt) > interval {
					fmt.Println("reap this entry from cache --> ", k)
					delete(c.cache, k)
				}
			}
			c.mutex.Unlock()
		}
	}
}
