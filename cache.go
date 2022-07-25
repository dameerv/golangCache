package cache

import "errors"

type Cache struct {
	cache map[string]interface{}
}

func New() *Cache {
	return &Cache{
		cache: make(map[string]interface{}),
	}
}
func (c *Cache) Set(key string, value interface{}) {
	c.cache[key] = value
}

func (c Cache) Get(key string) (interface{}, error) {
	if _, ok := c.cache[key]; !ok {
		return nil, errors.New("key not found")
	}
	return c.cache[key], nil
}

func (c *Cache) Delete(key string) error {
	if _, ok := c.cache[key]; !ok {
		return errors.New("can't delete: key not found")
	}
	delete(c.cache, key)
	return nil
}
