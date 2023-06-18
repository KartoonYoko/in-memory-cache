package internal

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache  map[string]*cacheValue
	locker *sync.RWMutex
}

func New() *Cache {
	cache := &Cache{
		cache: make(map[string]*cacheValue),
		locker: new(sync.RWMutex),
	}
	return cache
}

func (c *Cache) Set(key string, value any, ttl time.Duration) {
	c.locker.RLock()
	defer c.locker.RUnlock()

	c.cache[key] = &cacheValue{
		value:      value,
		expiryDate: time.Now().Add(ttl),
	}
}

func (c *Cache) Get(key string) (any, error) {
	c.locker.RLock()
	v, ok := c.cache[key]
	c.locker.RUnlock()

	if !ok {
		return nil, fmt.Errorf(keyNotFoundTemplate, ErrKeyNotFound, key)
	}

	if v.isTimeExpired() {
		c.Delete(key)
		return nil, fmt.Errorf(keyNotFoundTemplate, ErrKeyNotFound, key)
	}
	return v, nil
}

func (c *Cache) Delete(key string) {
	c.locker.RLock()
	defer c.locker.RUnlock()

	delete(c.cache, key)
}

// очистить все значения, у которых истекла дата существования
// func (c *Cache) cleanAllExpiredValues() {
// 	for key, cv := range c.cache {
// 		if cv.isTimeExpired() {
// 			fmt.Println("expired: ", cv)
// 			c.Delete(key)
// 		}
// 	}
// }
