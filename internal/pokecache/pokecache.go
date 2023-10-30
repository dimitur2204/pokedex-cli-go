package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	data      []byte
}

// Client -
type Cache struct {
	entries  map[string]cacheEntry
	interval time.Duration
	mux      *sync.Mutex
}

// NewClient -
func NewCache(interval time.Duration) Cache {
	cache := Cache{
		interval: interval,
		entries:  make(map[string]cacheEntry),
		mux:      &sync.Mutex{},
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	entry, ok := c.entries[key]
	c.mux.Unlock()
	return entry.data, ok
}

func (c *Cache) Set(key string, data []byte) {
	c.mux.Lock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		data:      data,
	}
	c.mux.Unlock()
}

func (c *Cache) Delete(key string) {
	c.mux.Lock()
	delete(c.entries, key)
	c.mux.Unlock()
}

func (c *Cache) reapLoop() {
	for {
		time.Sleep(c.interval)
		for k, e := range c.entries {
			if e.createdAt.Add(c.interval).Before(time.Now()) {
				c.Delete(k)
			}
		}
	}
}
