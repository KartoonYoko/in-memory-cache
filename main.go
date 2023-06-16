package main

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

func (c *Cache) Get(key string) (any, bool) {
	v, ok := c.cache[key]
	return v, ok
}

func (c *Cache) Delete(key string) {
	delete(c.cache, key)
}
