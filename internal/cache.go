package internal

import "fmt"

type Cache struct {
	cache map[string]any
}

func New() *Cache {
	return &Cache{
		cache: make(map[string]any),
	}
}

func (c *Cache) Set(key string, value any) {
	c.cache[key] = value
}

func (c *Cache) Get(key string) (any, error) {
	v, ok := c.cache[key]
	if !ok {
		return nil, fmt.Errorf(keyNotFoundTemplate, key)
	}
	return v, nil
}

func (c *Cache) Delete(key string) {
	delete(c.cache, key)
}
