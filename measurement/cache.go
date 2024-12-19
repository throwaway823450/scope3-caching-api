package measurement

import (
	"sync"
	"time"
)

type Cache struct {
	items map[string]CacheItem
	mutex sync.RWMutex
}

type CacheItem struct {
	Data      Row
	EntryTime time.Time
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]CacheItem),
		mutex: sync.RWMutex{},
	}
}

func (c *Cache) Set(key string, data Row) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.items[key] = CacheItem{
		Data:      data,
		EntryTime: time.Now(),
	}
}

func (c *Cache) Get(key string) (*Row, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	item, exists := c.items[key]
	if !exists {
		return nil, false
	}
	return &item.Data, true
}

func (c *Cache) GetWithTimestamp(key string) (CacheItem, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	item, exists := c.items[key]
	return item, exists
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.items, key)
}
