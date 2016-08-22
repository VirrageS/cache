package cache

import (
	"sync"
	"time"
)

type Cache struct {
	mutex sync.RWMutex
	ttl   time.Duration
	items map[string]interface{}
}

func (c *Cache) Set(key string, data interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.items[key] = data
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	return item, true
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.items, key)
}

func NewCache(ttl time.Duration) *Cache {
	cache := &Cache{
		ttl:   ttl,
		items: make(map[string]interface{}),
	}

	cache.startCleanupTimer()
	return cache
}

func (c *Cache) cleanup() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for key := range c.items {
		delete(c.items, key)
	}
}

func (c *Cache) startCleanupTimer() {
	ticker := time.Tick(c.ttl)
	go (func() {
		for {
			select {
			case <-ticker:
				c.cleanup()
			}
		}
	})()
}
