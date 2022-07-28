package cache

import (
	"errors"
	"sync"
)

type Cache struct {
	mu sync.RWMutex
	cache map[string]interface{}
}

func New() *Cache {
	return &Cache{
		cache: make(map[string]interface{}),
	}
}
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
}

func (c Cache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	if _, ok := c.cache[key]; !ok {
		return nil, errors.New("key not found")
	}

	defer c.mu.RUnlock()
	return c.cache[key], nil
}

func (c *Cache) Delete(key string) error {
	c.mu.RLock()
	
	if _, ok := c.cache[key]; !ok {
		return errors.New("can't delete: key not found")
	}

	defer c.mu.Unlock()
	delete(c.cache, key)
	return nil
}
