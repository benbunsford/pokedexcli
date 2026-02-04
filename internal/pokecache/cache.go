package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	mux      *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.cacheMap[key] = cacheEntry{time.Now(), val}
	c.mux.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	entry, ok := c.cacheMap[key]
	c.mux.Unlock()
	return entry.val, ok
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		for key, entry := range c.cacheMap {
			if time.Since(entry.createdAt) > interval {
				c.mux.Lock()
				delete(c.cacheMap, key)
				c.mux.Unlock()
			}
		}
		<-ticker.C
	}
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{}
	go cache.reapLoop(interval)
	return Cache{}
}
