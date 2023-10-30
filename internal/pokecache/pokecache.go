package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	data      *[]byte
}

// Client -
type Cache struct {
	entries  map[string]cacheEntry
	interval time.Duration
}

// NewClient -
func NewCache(interval time.Duration) Cache {
	cache := Cache{
		interval: interval,
		entries:  make(map[string]cacheEntry),
	}
	go cache.reapLoop(&sync.Mutex{})
	return cache
}

func (c *Cache) Get(key string, mux *sync.Mutex) (*[]byte, bool) {
	mux.Lock()
	entry, ok := c.entries[key]
	mux.Unlock()
	return entry.data, ok
}

func (c *Cache) Set(key string, data *[]byte, mux *sync.Mutex) {
	mux.Lock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		data:      data,
	}
	mux.Unlock()
}

func (c *Cache) Delete(key string, mux *sync.Mutex) {
	mux.Lock()
	delete(c.entries, key)
	mux.Unlock()
}

func (c *Cache) reapLoop(mux *sync.Mutex) {
	for {
		time.Sleep(c.interval)
		for k, e := range c.entries {
			if e.createdAt.Add(c.interval).Before(time.Now()) {
				c.Delete(k, mux)
			}
		}
	}
}
